// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Comment struct {
	ID        int64              `json:"id"`
	PostID    pgtype.Int8        `json:"post_id"`
	Text      string             `json:"text"`
	UserID    pgtype.Int8        `json:"user_id"`
	Status    string             `json:"status"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamp   `json:"updated_at"`
}

type Follow struct {
	FollowingUserID pgtype.Int8        `json:"following_user_id"`
	FollowedUserID  pgtype.Int8        `json:"followed_user_id"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
}

type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	// Content of the post
	Body      string             `json:"body"`
	UserID    pgtype.Int8        `json:"user_id"`
	Status    string             `json:"status"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamp   `json:"updated_at"`
}

type User struct {
	ID             int64              `json:"id"`
	Email          string             `json:"email"`
	HashedPassword string             `json:"hashed_password"`
	Username       string             `json:"username"`
	Role           string             `json:"role"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
}
