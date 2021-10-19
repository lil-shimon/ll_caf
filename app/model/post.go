package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
}

type Posts []Post
