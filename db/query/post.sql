-- name: CreatePost :one
INSERT INTO posts (
    title,
    body
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: GetPosts :many
SELECT * FROM posts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
SET title = $2, body = $3, status = $4, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;