package model

import "gorm.io/gorm"

/// Likes = 投稿にされたいいね
type Post struct {
	gorm.Model
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
	Likes   Likes
}

type Posts []Post
