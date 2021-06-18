package repository

import (
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/model"
)

func GetRepoTasks() (model.Tasks, error) {
	t := model.Tasks{}

	if err := database.DB.Find(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}
