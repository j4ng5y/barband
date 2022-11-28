package servercli

import (
	"os"
	"path/filepath"

	"github.com/j4ng5y/barband/pkg/cmd/cfg"
	"github.com/j4ng5y/barband/pkg/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/cobra"
)

func Run() error {
	var (
		c       = cfg.New()
		rootcmd = &cobra.Command{
			Use: "barband",
			Run: func(cmd *cobra.Command, args []string) {
				zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
				zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
				log.Logger = log.With().Caller().Logger()

				l, err := zerolog.ParseLevel(c.GetString("log.level"))
				if err != nil {
					log.Error().Err(err).Msg("invalid log level, defaulting to 'info'")
					l = zerolog.InfoLevel
				}
				zerolog.SetGlobalLevel(l)
				log.Debug().Msgf("set log level to %s", l.String())

				s, err := server.New(c)
				if err != nil {
					log.Fatal().Stack().Err(err).Send()
				}

				if err := s.Run(); err != nil {
					log.Fatal().Stack().Err(err).Send()
				}
			},
		}
	)

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	rootcmd.Flags().String("db.host", "localhost", "The hostname of the Database instance.")
	rootcmd.Flags().Uint("db.port", 5432, "The port of the Database instance.")
	rootcmd.Flags().String("db.sslmode", "disable", "The SSL Mode of the Database instance.")
	rootcmd.Flags().String("db.username", "barband", "The username of the Database.")
	rootcmd.Flags().String("db.password", "barband", "The password of the Database.")
	rootcmd.Flags().String("db.name", "barband", "The name of the Database.")
	rootcmd.Flags().String("db.migrationsDir", filepath.Join(pwd, "docs", "sql", "migrations"), "The absolute path to the migrations directory")

	rootcmd.Flags().String("server.addr", "0.0.0.0", "The IP Address to bind the gRPC server to.")
	rootcmd.Flags().Uint("server.port", 50051, "The TCP Port to bind the gRPC server to.")

	rootcmd.Flags().String("log.level", "info", "The log level to set. Available values: [trace, debug, info, warn, error, fatal, panic]")

	if err := c.BindPFlags(rootcmd.Flags()); err != nil {
		log.Fatal().Err(err).Send()
	}

	if err := rootcmd.Execute(); err != nil {
		log.Fatal().Err(err).Send()
	}

	return nil
}
