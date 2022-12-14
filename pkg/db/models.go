// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"

	"github.com/google/uuid"
)

type Band struct {
	ID        uuid.UUID
	Name      string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type BandMember struct {
	BandID      uuid.NullUUID
	ID          uuid.UUID
	Prefix      sql.NullString
	FirstName   string
	MiddleName  sql.NullString
	LastName    sql.NullString
	Suffix      sql.NullString
	NickName    sql.NullString
	PhoneNumber sql.NullInt32
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	DeletedAt   sql.NullTime
}

type Lyric struct {
	SongID uuid.NullUUID
	ID     uuid.UUID
	Part   sql.NullString
	Lyrics sql.NullString
}

type Set struct {
	SetListID uuid.NullUUID
	ID        uuid.UUID
	Name      sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type SetList struct {
	BandID    uuid.NullUUID
	ID        uuid.UUID
	Name      sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type Song struct {
	SetID                  uuid.NullUUID
	ID                     uuid.UUID
	Name                   sql.NullString
	CreatedAt              sql.NullTime
	UpdatedAt              sql.NullTime
	DeletedAt              sql.NullTime
	Title                  sql.NullString
	Genre                  sql.NullString
	Subgenre               sql.NullString
	PrimaryArtist          sql.NullString
	FeaturedArtists        sql.NullString
	Composer               sql.NullString
	Publisher              sql.NullString
	Producers              sql.NullString
	AdditionalContributors sql.NullString
	ExplicitContent        sql.NullBool
	LyricsLanguage         sql.NullString
	LyricsPublisher        sql.NullString
	CompositionOwner       sql.NullString
	YearOfComposition      sql.NullInt32
	MasterRecordingOwner   sql.NullString
	YearOfRecording        sql.NullInt32
	ReleaseLanguage        sql.NullString
}

type Tab struct {
	SongID uuid.NullUUID
	ID     uuid.UUID
	Part   sql.NullString
	Tabs   sql.NullString
}
