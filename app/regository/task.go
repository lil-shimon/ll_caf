package regository

import (
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/model"
)

func GetRepoTask() (model.Tasks, error) {
	t := model.Tasks{}

	if err := database.DB.Find(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}
