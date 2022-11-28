package server

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/j4ng5y/barband/pkg/db"
	"github.com/j4ng5y/barband/pkg/server/buffs"
)

func (s *Server) CreateBandMember(ctx context.Context, request *buffs.CreateBandMemberRequest) (*buffs.BandMember, error) {
	bid, err := uuid.Parse(request.GetBandId())
	if err != nil {
		return nil, err
	}
	id := uuid.New()

	member, err := s.db.CreateBandMember(ctx, db.CreateBandMemberParams{
		BandID:      uuid.NullUUID{UUID: bid},
		ID:          id,
		Prefix:      sql.NullString{String: request.GetPrefix()},
		FirstName:   request.GetFirstName(),
		MiddleName:  sql.NullString{String: request.GetMiddleName()},
		LastName:    sql.NullString{String: request.GetLastName()},
		Suffix:      sql.NullString{String: request.GetSuffix()},
		NickName:    sql.NullString{String: request.GetNickName()},
		PhoneNumber: sql.NullInt32{Int32: request.GetPhoneNumber()},
	})
	if err != nil {
		return nil, err
	}

	return &buffs.BandMember{
		BandId:      member.BandID.UUID.String(),
		Id:          member.ID.String(),
		Prefix:      member.Prefix.String,
		FirstName:   member.FirstName,
		MiddleName:  member.MiddleName.String,
		LastName:    member.LastName.String,
		Suffix:      member.Suffix.String,
		NickName:    member.NickName.String,
		PhoneNumber: member.PhoneNumber.Int32,
	}, nil
}

func (s *Server) ReadBandMember(ctx context.Context, request *buffs.ReadBandMemberRequest) (*buffs.BandMember, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}
	member, err := s.db.ReadBandMember(ctx, id)
	if err != nil {
		return nil, err
	}
	return &buffs.BandMember{
		BandId:      member.BandID.UUID.String(),
		Id:          member.ID.String(),
		Prefix:      member.Prefix.String,
		FirstName:   member.FirstName,
		MiddleName:  member.MiddleName.String,
		LastName:    member.LastName.String,
		Suffix:      member.Suffix.String,
		NickName:    member.NickName.String,
		PhoneNumber: member.PhoneNumber.Int32,
	}, nil
}

func (s *Server) UpdateBandMember(ctx context.Context, request *buffs.UpdateBandMemberRequest) (*buffs.BandMember, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}
	if err := s.db.UpdateBandMember(ctx, db.UpdateBandMemberParams{
		ID:          id,
		Prefix:      sql.NullString{String: request.GetPrefix()},
		FirstName:   request.GetFirstName(),
		MiddleName:  sql.NullString{String: request.GetMiddleName()},
		LastName:    sql.NullString{String: request.GetLastName()},
		Suffix:      sql.NullString{String: request.GetSuffix()},
		NickName:    sql.NullString{String: request.GetNickName()},
		PhoneNumber: sql.NullInt32{Int32: request.GetPhoneNumber()},
	}); err != nil {
		return nil, err
	}
	member, err := s.db.ReadBandMember(ctx, id)
	if err != nil {
		return nil, err
	}
	return &buffs.BandMember{
		BandId:      member.BandID.UUID.String(),
		Id:          member.ID.String(),
		Prefix:      member.Prefix.String,
		FirstName:   member.FirstName,
		MiddleName:  member.MiddleName.String,
		LastName:    member.LastName.String,
		Suffix:      member.Suffix.String,
		NickName:    member.NickName.String,
		PhoneNumber: member.PhoneNumber.Int32,
	}, nil
}

func (s *Server) DeleteBandMember(ctx context.Context, request *buffs.DeleteBandMemberRequest) (*buffs.Empty, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}

	if err := s.db.DeleteBandMember(ctx, id); err != nil {
		return nil, err
	}

	return &buffs.Empty{}, nil
}
