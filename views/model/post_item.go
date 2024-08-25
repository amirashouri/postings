package model

import db "main/db/sqlc"

type PostItem struct {
	Post       db.Post
	LikesCount int
	IsLiked    bool
}
