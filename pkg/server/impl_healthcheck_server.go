package server

import (
	"context"

	"github.com/j4ng5y/barband/pkg/server/buffs"
)

func (s *Server) HealthCheck(ctx context.Context, request *buffs.Empty) (*buffs.Empty, error) {
	return &buffs.Empty{}, nil
}
