package server

import (
	"context"

	"github.com/j4ng5y/barband/pkg/server/buffs"
)

func (s *Server) CreateSong(ctx context.Context, request *buffs.CreateSongRequest) (*buffs.Song, error) {
	return nil, nil
}

func (s *Server) ReadSong(ctx context.Context, request *buffs.ReadSongRequest) (*buffs.Song, error) {
	return nil, nil
}

func (s *Server) UpdateSong(ctx context.Context, request *buffs.UpdateSongRequest) (*buffs.Song, error) {
	return nil, nil
}

func (s *Server) DeleteSong(ctx context.Context, request *buffs.DeleteSongRequest) (*buffs.Empty, error) {
	return nil, nil
}

func (s *Server) AddLyrics(ctx context.Context, request *buffs.AddLyricsRequest) (*buffs.Lyric, error) {
	return nil, nil
}

func (s *Server) AddTabs(ctx context.Context, request *buffs.AddTabsRequest) (*buffs.Tab, error) {
	return nil, nil
}
