-- name: CreateBandMember :one
INSERT INTO band_members (
    band_id, id, prefix, first_name, middle_name, last_name, suffix, nick_name, phone_number, created_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9, current_timestamp
         ) RETURNING *;

-- name: ReadBandMember :one
SELECT * FROM band_members
WHERE id = $1 LIMIT 1;

-- name: ReadBandMembers :many
SELECT * FROM band_members;

-- name: UpdateBandMember :exec
UPDATE band_members
SET prefix = $2,
    first_name = $3,
    middle_name = $4,
    last_name = $5,
    suffix = $6,
    nick_name = $7,
    phone_number = $8,
    updated_at = current_timestamp
WHERE id = $1;

-- name: DeleteBandMember :exec
UPDATE band_members SET deleted_at = current_timestamp
WHERE id = $1;