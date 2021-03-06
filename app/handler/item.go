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

func ShowItem(c echo.Context) error {
  id := c.Param("id")
  i, _ := service.ShowItem(id)
  return c.JSON(http.StatusOK, i)
}

func UpdateItem(c echo.Context) error {
  i, _ := service.ShowItem(c.Param("id"))
  c.Bind(&i)
  i, _ = service.UpdateItem(i)
  return c.JSON(http.StatusOK, i)
}

func GetItemsTypeId(c echo.Context) error {
  is, err := service.GetItemsByTypeId(c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, err)
  }
  return c.JSON(http.StatusOK, is)
}
