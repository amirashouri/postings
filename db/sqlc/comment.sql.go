// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: comment.sql

package db

import (
	"context"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comments (
    text
) VALUES (
    $1
) RETURNING id, post_id, text, user_id, status, created_at, updated_at
`

func (q *Queries) CreateComment(ctx context.Context, text string) (Comment, error) {
	row := q.db.QueryRow(ctx, createComment, text)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.Text,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteComment, id)
	return err
}

const getComment = `-- name: GetComment :one
SELECT id, post_id, text, user_id, status, created_at, updated_at FROM comments
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetComment(ctx context.Context, id int64) (Comment, error) {
	row := q.db.QueryRow(ctx, getComment, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.Text,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getComments = `-- name: GetComments :many
SELECT id, post_id, text, user_id, status, created_at, updated_at FROM comments
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetCommentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetComments(ctx context.Context, arg GetCommentsParams) ([]Comment, error) {
	rows, err := q.db.Query(ctx, getComments, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.Text,
			&i.UserID,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateComment = `-- name: UpdateComment :one
UPDATE comments
SET text = $2, status = $3, updated_at = now()
WHERE id = $1
RETURNING id, post_id, text, user_id, status, created_at, updated_at
`

type UpdateCommentParams struct {
	ID     int64  `json:"id"`
	Text   string `json:"text"`
	Status string `json:"status"`
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error) {
	row := q.db.QueryRow(ctx, updateComment, arg.ID, arg.Text, arg.Status)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.Text,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
