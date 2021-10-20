package repository

import (
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/model"
)

/// 投稿に紐づけられたいいねを取得
func GetLikesByPostId(pid string) (model.Likes, error) {
	p := model.Likes{}
	if err := database.DB.Where("post_id = ?", pid).Where("type = ?", 1).Find(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}
