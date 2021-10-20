package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/service"
)

func GetUser(c echo.Context) error {
	u, _ := service.GetUser(c.Param("id"))
	return c.JSON(http.StatusOK, u)
}

func StoreUser(c echo.Context) error {
	u := model.User{}
	nu, _ := service.StoreUser(u)
	return c.JSON(http.StatusOK, nu)
}
