-- name: CreateSong :one
INSERT INTO songs (
    set_id, id, title, genre, subgenre, primary_artist, featured_artists, composer, publisher, producers, additional_contributors, explicit_content, lyrics_language, lyrics_publisher, composition_owner, year_of_composition, master_recording_owner, year_of_recording, release_language, created_at
) VALUES (
             $1, $2, $3 ,$4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, current_timestamp
         ) RETURNING *;

-- name: ReadSong :one
SELECT * FROM songs
WHERE id = $1 LIMIT 1;

-- name: ReadSongs :many
SELECT * FROM songs;

-- name: UpdateSong :exec
UPDATE songs
SET title = $2,
    genre = $3,
    subgenre = $4,
    primary_artist = $5,
    featured_artists = $6,
    composer = $7,
    publisher = $8,
    producers = $9,
    additional_contributors = $10,
    explicit_content = $11,
    lyrics_language = $12,
    lyrics_publisher = $13,
    composition_owner = $14,
    year_of_composition = $15,
    master_recording_owner = $16,
    year_of_recording = $17,
    release_language = $18,
    updated_at = current_timestamp
WHERE id = $1;

-- name: DeleteSong :exec
UPDATE songs SET deleted_at = current_timestamp
WHERE id = $1;

-- name: AddLyric :one
INSERT INTO lyrics (
    song_id, id, part, lyrics, created_at
) VALUES (
    $1, $2, $3, $3, current_timestamp
) RETURNING *;

-- name: AddTab :one
INSERT INTO tabs (
    song_id, id, part, tabs, created_at
) VALUES (
     $1, $2, $3, $3, current_timestamp
 ) RETURNING *;