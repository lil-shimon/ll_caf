package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/service"
)

func GetItems(c echo.Context) error {
  i, _ := service.GetItems()
  return c.JSON(http.StatusOK, i)
}

func StoreItem(c echo.Context) error {
  i := new(model.Item)
  c.Bind(&i)
  ni, _ := service.StoreItem(*i)
  return c.JSON(http.StatusOK, ni)
}
