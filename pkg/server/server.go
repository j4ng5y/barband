package server

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/j4ng5y/barband/pkg/db"
	"github.com/j4ng5y/barband/pkg/server/buffs"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type (
	Server struct {
		cfg *viper.Viper
		db  *db.Queries
		buffs.UnimplementedBandServiceServer
		buffs.UnimplementedHealthcheckServiceServer
		buffs.UnimplementedBandMemberServiceServer
		buffs.UnimplementedSetListServiceServer
		buffs.UnimplementedSetServiceServer
		buffs.UnimplementedSongServiceServer
	}
)

func New(cfg *viper.Viper) (*Server, error) {
	database, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.GetString("db.host"),
		cfg.GetUint("db.port"),
		cfg.GetString("db.username"),
		cfg.GetString("db.password"),
		cfg.GetString("db.name"),
		cfg.GetString("db.sslmode"),
	))
	if err != nil {
		return nil, err
	}

	drv, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", cfg.GetString("db.migrationsDir")), "postgres", drv)
	if err != nil {
		return nil, err
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {

			log.Info().Msg("no migrations to run")
		} else {
			if err := m.Down(); err != nil {
				log.Fatal().Err(err).Send()
			}
			return nil, err
		}
	}
	v, d, err := m.Version()
	if err != nil {
		log.Error().Err(err).Send()
	}
	log.Info().Uint("version", v).Bool("dirty", d).Msg("migrations successfully ran")

	s := &Server{
		cfg: cfg,
		db:  db.New(database),
	}

	return s, nil
}

func (s *Server) Run() error {
	runChan := make(chan os.Signal, 1)

	grpcAddr := fmt.Sprintf("%s:%d", s.cfg.GetString("server.addr"), s.cfg.GetUint("server.port"))
	log.Info().Msgf("Running gRPC server on: %s", grpcAddr)

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return err
	}

	opts := []logging.Option{
		logging.WithLevels(logging.DefaultServerCodeToLevel),
		logging.WithDecider(func(methodFullName string, err error) logging.Decision {
			switch methodFullName {
			case "/BarBand/HealthCheck":
				return logging.NoLogCall
			case "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo":
				return logging.NoLogCall
			}
			return logging.LogStartAndFinishCall
		}),
	}

	S := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(),
			logging.UnaryServerInterceptor(
				grpczerolog.InterceptorLogger(log.Logger), opts...,
			),
		),
		grpc.ChainStreamInterceptor(
			recovery.StreamServerInterceptor(),
			logging.StreamServerInterceptor(
				grpczerolog.InterceptorLogger(log.Logger), opts...,
			),
		),
	)
	buffs.RegisterBandServiceServer(S, s)
	buffs.RegisterBandMemberServiceServer(S, s)
	buffs.RegisterHealthcheckServiceServer(S, s)
	buffs.RegisterSetServiceServer(S, s)
	buffs.RegisterSetListServiceServer(S, s)
	buffs.RegisterSongServiceServer(S, s)
	reflection.Register(S)

	go func() {
		if err := S.Serve(lis); err != nil {
			log.Fatal().Err(err).Send()
		}
	}()

	sig := <-runChan

	log.Info().Msgf("Stopping gRPC server due to: %s", sig.String())
	S.GracefulStop()

	return nil
}
