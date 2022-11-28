package server

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/j4ng5y/barband/pkg/db"
	"github.com/j4ng5y/barband/pkg/server/buffs"
)

func (s *Server) CreateSet(ctx context.Context, request *buffs.CreateSetRequest) (*buffs.Set, error) {
	slid, err := uuid.Parse(request.GetSetlistId())
	if err != nil {
		return nil, err
	}
	id := uuid.New()

	set, err := s.db.CreateSet(ctx, db.CreateSetParams{
		SetListID: uuid.NullUUID{UUID: slid},
		ID:        id,
		Name:      sql.NullString{String: request.GetName()},
	})
	if err != nil {
		return nil, err
	}

	return &buffs.Set{
		SetlistId: set.SetListID.UUID.String(),
		Id:        set.ID.String(),
		Name:      set.Name.String,
	}, nil
}

func (s *Server) ReadSet(ctx context.Context, request *buffs.ReadSetRequest) (*buffs.Set, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}

	set, err := s.db.ReadSet(ctx, id)
	if err != nil {
		return nil, err
	}

	return &buffs.Set{
		SetlistId: set.SetListID.UUID.String(),
		Id:        set.ID.String(),
		Name:      set.Name.String,
	}, nil
}

func (s *Server) UpdateSet(ctx context.Context, request *buffs.UpdateSetRequest) (*buffs.Set, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}

	if err := s.db.UpdateSet(ctx, db.UpdateSetParams{
		ID:   id,
		Name: sql.NullString{String: request.GetName()},
	}); err != nil {
		return nil, err
	}

	set, err := s.db.ReadSet(ctx, id)
	if err != nil {
		return nil, err
	}

	return &buffs.Set{
		SetlistId: set.SetListID.UUID.String(),
		Id:        set.ID.String(),
		Name:      set.Name.String,
	}, nil
}

func (s *Server) DeleteSet(ctx context.Context, request *buffs.DeleteSetRequest) (*buffs.Empty, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}
	if err := s.db.DeleteSet(ctx, id); err != nil {
		return nil, err
	}
	return &buffs.Empty{}, nil
}
