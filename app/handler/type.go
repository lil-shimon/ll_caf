package handler

import (
	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/model"
	"net/http"
)

func CreateType(c echo.Context) error {
	t := model.Types{}
	if err := c.Bind(&t); err != nil {
		return err
	}
	database.DB.Create(&t)
	return c.JSON(http.StatusCreated, t)
}

func GetType(c echo.Context) error {
	t := new(model.Types)
	if err := c.Bind(t); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, t)
}
