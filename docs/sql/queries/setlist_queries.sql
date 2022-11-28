-- name: CreateSetList :one
INSERT INTO set_lists (
    band_id, id, name, created_at
) VALUES (
             $1, $2, $3, current_timestamp
         ) RETURNING *;

-- name: ReadSetList :one
SELECT * FROM set_lists
WHERE id = $1 LIMIT 1;

-- name: ReadSetLists :many
SELECT * FROM set_lists;

-- name: UpdateSetList :exec
UPDATE set_lists SET name = $2, updated_at = current_timestamp
WHERE id = $1;

-- name: DeleteSetList :exec
UPDATE set_lists SET deleted_at = current_timestamp
WHERE id = $1;