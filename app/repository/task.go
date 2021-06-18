package repository

import (
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/model"
)

// GetRepoTasks
// return [type] Tasks, error
//
// handler.GetTasks
///**
func GetRepoTasks() (model.Tasks, error) {
	t := model.Tasks{}

	if err := database.DB.Find(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}

// ShowRepoTask
// param id string
// return [type] Task, error
//
// handler.ShowTask
///**
func ShowRepoTask(id string) (model.Task, error) {
	t := model.Task{}

	if err := database.DB.Where("id = ?", id).First(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}

func UpdateRepoTask(t model.Task) (model.Task, error) {
	if err := database.DB.Save(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}