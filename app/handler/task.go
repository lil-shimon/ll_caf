package handler

import (
	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/repository"
	"net/http"
)

func GetTasks(c echo.Context) error {
	t, _ := repository.GetRepoTasks()
	return c.JSON(http.StatusOK, t)
}

func ShowTask()