package model

import db "main/db/sqlc"

type PostItem struct {
	UserName   string
	Post       db.Post
	LikesCount int
	IsLiked    bool
	ShowDelete bool
}
