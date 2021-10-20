package model

import "time"

type Like struct {
	Id        uint      `json:"id"`
	UserId    int       `json:"user_id"`
	PostId    int       `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Likes []*Like
