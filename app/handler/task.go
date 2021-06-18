package handler

import (
	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/repository"
	"net/http"
)

func GetTasks() error {
	t, _ := repository.GetRepoTasks()
	return c.JSON(http.StatusOK, t)
}
