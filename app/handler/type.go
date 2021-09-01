package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/repository"
)

func GetTypes(c echo.Context) error {
  t, _ := repository.GetTypes()
  return c.JSON(http.StatusOK, t)
}

func ShowType (c echo.Context) error {
  id := c.Param("id")
  t, _ := repository.ShowType(id)
  return c.JSON(http.StatusOK, t)
}

func StoreType (c echo.Context) error {
  t := new(model.Type)
  c.Bind(t)
  nt, _ := repository.StoreType(*t)
  return c.JSON(http.StatusOK, nt)
}

func UpdateType (c echo.Context) error {
  t, _ := repository.ShowType(c.Param("id"))
  if err := c.Bind(&t); err != nil {
    return err
  }
  t, _ = repository.UpdateType(t)
  return c.JSON(http.StatusOK, t)
}
