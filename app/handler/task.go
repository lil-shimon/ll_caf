package handler

import (
	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/repository"
)

func GetTasks() error {
	t, _ := repository.GetRepoTasks()
	return t
}
