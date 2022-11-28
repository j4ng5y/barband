-- name: CreateSet :one
INSERT INTO sets (
    set_list_id, id, name, created_at
) VALUES (
             $1, $2, $3, current_timestamp
         ) RETURNING *;

-- name: ReadSet :one
SELECT * FROM sets
WHERE id = $1 LIMIT 1;

-- name: ReadSets :many
SELECT * FROM sets;

-- name: UpdateSet :exec
UPDATE sets SET name = $2, updated_at = current_timestamp
WHERE id = $1;

-- name: DeleteSet :exec
UPDATE sets SET deleted_at = current_timestamp
WHERE id = $1;