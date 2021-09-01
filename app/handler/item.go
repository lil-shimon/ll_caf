package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/service"
)

func GetItems(c echo.Context) error {
  i, _ := service.GetItems()
  return c.JSON(http.StatusOK, i)
}
