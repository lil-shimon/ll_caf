package repository

import (
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/model"
)

/// ユーザーごとの投稿をとる
func GetPostsByUserId(uid string) (model.Posts, error) {
	p := model.Posts{}
	if err := database.DB.Where("user_id = ?", uid).Find(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}
