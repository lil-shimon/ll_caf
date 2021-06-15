package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/model"
	"net/http"
)

func CreateType (c echo.Context) error {
	t := new(model.TaskType)
	if err := c.Bind(t); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, t)
}

func GetType (c echo.Context) error {
	t := new(model.TaskType)
	if err := c.Bind(t); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, t)
}