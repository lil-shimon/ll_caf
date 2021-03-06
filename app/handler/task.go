package handler

import (
	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/repository"
	"net/http"
)

// GetTasks
// param c [echo.Context]
// return json [model.Tasks], error
//
// repository.GetRepoTasks
///**
func GetTasks(c echo.Context) error {
	t, _ := repository.GetRepoTasks()
	return c.JSON(http.StatusOK, t)
}

// ShowTask
// param c [echo.Context]
// return json, error
//
// repository.ShowRepoTask
///**
func ShowTask(c echo.Context) error {
	id := c.Param("id")
	t, _ := repository.ShowRepoCategory(string(id))
	return c.JSON(http.StatusOK, t)
}

// UpdateTask
// param c echo.Context
// return t model.Task
//
// repository.UpdateRepoTask
///**
func UpdateTask(c echo.Context) error {
	t := model.Task{}
	if err := c.Bind(&t); err != nil {
		return err
	}
	t, _ = repository.UpdateRepoTask(t)
	return c.JSON(http.StatusOK, t)
}

func DeleteTask(c echo.Context) error {
	t := model.Task{}
	if err := c.Bind(&t); err != nil {
		return err
	}
	database.DB.Delete(&t)
	return c.JSON(http.StatusNoContent, t)
}