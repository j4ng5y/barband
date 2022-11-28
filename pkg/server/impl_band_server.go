package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/j4ng5y/barband/pkg/db"
	"github.com/j4ng5y/barband/pkg/server/buffs"
	"github.com/rs/zerolog/log"
)

func (s *Server) CreateBand(ctx context.Context, request *buffs.CreateBandRequest) (*buffs.Band, error) {
	band, err := s.db.CreateBand(ctx, db.CreateBandParams{
		ID:   uuid.New(),
		Name: request.Name,
	})
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	return &buffs.Band{
		Id:   band.ID.String(),
		Name: band.Name,
	}, nil
}

func (s *Server) ReadBand(ctx context.Context, request *buffs.ReadBandRequest) (*buffs.Band, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}
	band, err := s.db.ReadBand(ctx, id)
	if err != nil {
		return nil, err
	}
	return &buffs.Band{Id: band.ID.String(), Name: band.Name}, nil
}

func (s *Server) UpdateBand(ctx context.Context, request *buffs.UpdateBandRequest) (*buffs.Band, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}
	if err := s.db.UpdateBand(ctx, db.UpdateBandParams{
		ID:   id,
		Name: request.GetName(),
	}); err != nil {
		return nil, err
	}
	band, err := s.db.ReadBand(ctx, id)
	if err != nil {
		return nil, err
	}
	return &buffs.Band{
		Id:   band.ID.String(),
		Name: band.Name,
	}, nil
}

func (s *Server) DeleteBand(ctx context.Context, request *buffs.DeleteBandRequest) (*buffs.Empty, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}
	if err := s.db.DeleteBand(ctx, id); err != nil {
		return nil, err
	}
	return &buffs.Empty{}, nil
}
