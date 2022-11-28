-- name: CreateBand :one
INSERT INTO bands (
    id, name, created_at
) VALUES (
    $1, $2, current_timestamp
) RETURNING *;

-- name: ReadBand :one
SELECT * FROM bands
WHERE id = $1 LIMIT 1;

-- name: ReadBands :many
SELECT * FROM bands;

-- name: UpdateBand :exec
UPDATE bands SET name = $2, updated_at = current_timestamp
WHERE id = $1;

-- name: DeleteBand :exec
UPDATE bands SET deleted_at = current_timestamp
WHERE id = $1;