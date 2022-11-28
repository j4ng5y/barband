package server

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/j4ng5y/barband/pkg/db"
	"github.com/j4ng5y/barband/pkg/server/buffs"
)

func (s *Server) CreateSong(ctx context.Context, request *buffs.CreateSongRequest) (*buffs.Song, error) {
	setid, err := uuid.Parse(request.GetSetId())
	if err != nil {
		return nil, err
	}
	id := uuid.New()
	var title string
	if request.GetName() != "" {
		title = request.GetName()
	} else {
		title = request.GetTitle()
	}
	song, err := s.db.CreateSong(ctx, db.CreateSongParams{
		SetID:                  uuid.NullUUID{UUID: setid},
		ID:                     id,
		Title:                  sql.NullString{String: title},
		Genre:                  sql.NullString{String: request.GetGenre()},
		Subgenre:               sql.NullString{String: request.GetSubgenre()},
		PrimaryArtist:          sql.NullString{String: request.GetPrimaryArtist()},
		FeaturedArtists:        sql.NullString{String: request.GetFeaturedArtists()},
		Composer:               sql.NullString{String: request.GetComposer()},
		Publisher:              sql.NullString{String: request.GetPublisher()},
		Producers:              sql.NullString{String: request.GetProducers()},
		AdditionalContributors: sql.NullString{String: request.GetAdditionalContributors()},
		ExplicitContent:        sql.NullBool{Bool: request.GetExplicitContent()},
		LyricsLanguage:         sql.NullString{String: request.GetLyricsLanguage()},
		LyricsPublisher:        sql.NullString{String: request.GetLyricsPublisher()},
		CompositionOwner:       sql.NullString{String: request.GetCompositionOwner()},
		YearOfComposition:      sql.NullInt32{Int32: request.GetYearOfComposition()},
		MasterRecordingOwner:   sql.NullString{String: request.GetMasterRecordingOwner()},
		YearOfRecording:        sql.NullInt32{Int32: request.GetYearOfRecording()},
		ReleaseLanguage:        sql.NullString{String: request.GetReleaseLanguage()},
	})
	if err != nil {
		return nil, err
	}

	return &buffs.Song{
		SetId:                  song.SetID.UUID.String(),
		Id:                     song.ID.String(),
		NameOrTitle:            &buffs.Song_Title{Title: song.Title.String},
		Genre:                  song.Genre.String,
		Subgenre:               song.Subgenre.String,
		PrimaryArtist:          song.PrimaryArtist.String,
		FeaturedArtists:        song.FeaturedArtists.String,
		Composer:               song.Composer.String,
		Publisher:              song.Publisher.String,
		Producers:              song.Producers.String,
		AdditionalContributors: song.AdditionalContributors.String,
		ExplicitContent:        song.ExplicitContent.Bool,
		LyricsLanguage:         song.LyricsLanguage.String,
		LyricsPublisher:        song.LyricsPublisher.String,
		CompositionOwner:       song.CompositionOwner.String,
		YearOfComposition:      song.YearOfComposition.Int32,
		MasterRecordingOwner:   song.MasterRecordingOwner.String,
		YearOfRecording:        song.YearOfRecording.Int32,
		ReleaseLanguage:        song.ReleaseLanguage.String,
	}, nil
}

func (s *Server) ReadSong(ctx context.Context, request *buffs.ReadSongRequest) (*buffs.Song, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}
	song, err := s.db.ReadSong(ctx, id)
	if err != nil {
		return nil, err
	}
	return &buffs.Song{
		SetId:                  song.SetID.UUID.String(),
		Id:                     song.ID.String(),
		NameOrTitle:            &buffs.Song_Title{Title: song.Title.String},
		Genre:                  song.Genre.String,
		Subgenre:               song.Subgenre.String,
		PrimaryArtist:          song.PrimaryArtist.String,
		FeaturedArtists:        song.FeaturedArtists.String,
		Composer:               song.Composer.String,
		Publisher:              song.Publisher.String,
		Producers:              song.Producers.String,
		AdditionalContributors: song.AdditionalContributors.String,
		ExplicitContent:        song.ExplicitContent.Bool,
		LyricsLanguage:         song.LyricsLanguage.String,
		LyricsPublisher:        song.LyricsPublisher.String,
		CompositionOwner:       song.CompositionOwner.String,
		YearOfComposition:      song.YearOfComposition.Int32,
		MasterRecordingOwner:   song.MasterRecordingOwner.String,
		YearOfRecording:        song.YearOfRecording.Int32,
		ReleaseLanguage:        song.ReleaseLanguage.String,
	}, nil
}

func (s *Server) UpdateSong(ctx context.Context, request *buffs.UpdateSongRequest) (*buffs.Song, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}
	var title string
	if request.GetName() != "" {
		title = request.GetName()
	} else {
		title = request.GetTitle()
	}
	if err := s.db.UpdateSong(ctx, db.UpdateSongParams{
		ID:                     id,
		Title:                  sql.NullString{String: title},
		Genre:                  sql.NullString{String: request.GetGenre()},
		Subgenre:               sql.NullString{String: request.GetSubgenre()},
		PrimaryArtist:          sql.NullString{String: request.GetPrimaryArtist()},
		FeaturedArtists:        sql.NullString{String: request.GetFeaturedArtists()},
		Composer:               sql.NullString{String: request.GetComposer()},
		Publisher:              sql.NullString{String: request.GetPublisher()},
		Producers:              sql.NullString{String: request.GetProducers()},
		AdditionalContributors: sql.NullString{String: request.GetAdditionalContributors()},
		ExplicitContent:        sql.NullBool{Bool: request.GetExplicitContent()},
		LyricsLanguage:         sql.NullString{String: request.GetLyricsLanguage()},
		LyricsPublisher:        sql.NullString{String: request.GetLyricsPublisher()},
		CompositionOwner:       sql.NullString{String: request.GetCompositionOwner()},
		YearOfComposition:      sql.NullInt32{Int32: request.GetYearOfComposition()},
		MasterRecordingOwner:   sql.NullString{String: request.GetMasterRecordingOwner()},
		YearOfRecording:        sql.NullInt32{Int32: request.GetYearOfRecording()},
		ReleaseLanguage:        sql.NullString{String: request.GetReleaseLanguage()},
	}); err != nil {
		return nil, err
	}

	song, err := s.db.ReadSong(ctx, id)
	if err != nil {
		return nil, err
	}

	return &buffs.Song{
		SetId:                  song.SetID.UUID.String(),
		Id:                     song.ID.String(),
		NameOrTitle:            &buffs.Song_Title{Title: song.Title.String},
		Genre:                  song.Genre.String,
		Subgenre:               song.Subgenre.String,
		PrimaryArtist:          song.PrimaryArtist.String,
		FeaturedArtists:        song.FeaturedArtists.String,
		Composer:               song.Composer.String,
		Publisher:              song.Publisher.String,
		Producers:              song.Producers.String,
		AdditionalContributors: song.AdditionalContributors.String,
		ExplicitContent:        song.ExplicitContent.Bool,
		LyricsLanguage:         song.LyricsLanguage.String,
		LyricsPublisher:        song.LyricsPublisher.String,
		CompositionOwner:       song.CompositionOwner.String,
		YearOfComposition:      song.YearOfComposition.Int32,
		MasterRecordingOwner:   song.MasterRecordingOwner.String,
		YearOfRecording:        song.YearOfRecording.Int32,
		ReleaseLanguage:        song.ReleaseLanguage.String,
	}, nil
}

func (s *Server) DeleteSong(ctx context.Context, request *buffs.DeleteSongRequest) (*buffs.Empty, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}
	if err := s.db.DeleteSong(ctx, id); err != nil {
		return nil, err
	}
	return &buffs.Empty{}, nil
}

func (s *Server) AddLyrics(ctx context.Context, request *buffs.AddLyricsRequest) (*buffs.Lyric, error) {
	sid, err := uuid.Parse(request.GetSongId())
	if err != nil {
		return nil, err
	}
	id := uuid.New()

	l, err := s.db.AddLyric(ctx, db.AddLyricParams{
		SongID: uuid.NullUUID{UUID: sid},
		ID:     id,
		Part:   sql.NullString{String: request.GetPart()},
		Lyrics: sql.NullString{String: request.GetLyrics()},
	})
	if err != nil {
		return nil, err
	}
	return &buffs.Lyric{
		SongId: l.SongID.UUID.String(),
		Id:     l.ID.String(),
		Part:   l.Part.String,
		Lyrics: l.Lyrics.String,
	}, nil
}

func (s *Server) AddTabs(ctx context.Context, request *buffs.AddTabsRequest) (*buffs.Tab, error) {
	sid, err := uuid.Parse(request.GetSongId())
	if err != nil {
		return nil, err
	}
	id := uuid.New()

	t, err := s.db.AddTab(ctx, db.AddTabParams{
		SongID: uuid.NullUUID{UUID: sid},
		ID:     id,
		Part:   sql.NullString{String: request.GetPart()},
		Tabs:   sql.NullString{String: request.GetTabs()},
	})
	if err != nil {
		return nil, err
	}
	return &buffs.Tab{
		SongId: t.SongID.UUID.String(),
		Id:     t.ID.String(),
		Part:   t.Part.String,
		Tabs:   t.Tabs.String,
	}, nil
}
