-- name: CreateLike :exec
INSERT INTO likes (
    user_id,
    post_id
) VALUES (
    $1, $2
);

-- name: GetLikes :many
SELECT * FROM likes
WHERE post_id = $1;

-- name: DeleteLike :exec
DELETE FROM likes
WHERE id = $1;