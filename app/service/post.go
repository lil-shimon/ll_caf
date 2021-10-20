package service

import (
	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/repository"
)

/// ユーザーごとに投稿をとる
func GetPostsByUserId(uid string) (model.Posts, error) {
	ps, _ := repository.GetPostsByUserId(uid)

	/// 投稿にいいね情報を付与
	for _, p := range ps {
		p.Likes, _ = repository.GetLikesByPostId(p.ID)
	}

	return ps, nil
}
