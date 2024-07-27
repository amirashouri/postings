-- name: CreateComment :one
INSERT INTO comments (
    text
) VALUES (
    $1
) RETURNING *;

-- name: GetComment :one
SELECT * FROM comments
WHERE id = $1 LIMIT 1;

-- name: GetComments :many
SELECT * FROM comments
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateComment :one
UPDATE comments
SET text = $2, status = $3, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;