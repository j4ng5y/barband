package server

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/j4ng5y/barband/pkg/db"
	"github.com/j4ng5y/barband/pkg/server/buffs"
)

func (s *Server) CreateSetList(ctx context.Context, request *buffs.CreateSetListRequest) (*buffs.SetList, error) {
	bid, err := uuid.Parse(request.GetBandId())
	if err != nil {
		return nil, err
	}
	id := uuid.New()

	setlist, err := s.db.CreateSetList(ctx, db.CreateSetListParams{
		BandID: uuid.NullUUID{UUID: bid},
		ID:     id,
		Name:   sql.NullString{String: request.GetName()},
	})
	if err != nil {
		return nil, err
	}

	return &buffs.SetList{
		BandId: setlist.BandID.UUID.String(),
		Id:     setlist.ID.String(),
		Name:   setlist.Name.String,
	}, nil
}

func (s *Server) ReadSetList(ctx context.Context, request *buffs.ReadSetListRequest) (*buffs.SetList, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}

	setlist, err := s.db.ReadSetList(ctx, id)
	if err != nil {
		return nil, err
	}

	return &buffs.SetList{
		BandId: setlist.BandID.UUID.String(),
		Id:     setlist.ID.String(),
		Name:   setlist.Name.String,
	}, nil
}

func (s *Server) UpdateSetList(ctx context.Context, request *buffs.UpdateSetListRequest) (*buffs.SetList, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}

	if err := s.db.UpdateSetList(ctx, db.UpdateSetListParams{
		ID:   id,
		Name: sql.NullString{String: request.GetName()},
	}); err != nil {
		return nil, err
	}

	setlist, err := s.db.ReadSetList(ctx, id)
	if err != nil {
		return nil, err
	}

	return &buffs.SetList{
		BandId: setlist.BandID.UUID.String(),
		Id:     setlist.ID.String(),
		Name:   setlist.Name.String,
	}, nil
}

func (s *Server) DeleteSetList(ctx context.Context, request *buffs.DeleteSetListRequest) (*buffs.Empty, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}
	if err := s.db.DeleteSetList(ctx, id); err != nil {
		return nil, err
	}
	return &buffs.Empty{}, nil
}
