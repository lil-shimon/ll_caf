package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/service"
)

func GetLogs(c echo.Context) error {
	ls, _ := service.GetLogs()
	return c.JSON(http.StatusOK, ls)
}

func GetLog(c echo.Context) error {
	id := c.Param("id")
	l, _ := service.GetLog(id)
	return c.JSON(http.StatusOK, l)
}

func StoreLog(c echo.Context) error {
	l := new(model.Log)
	if err := c.Bind(&l); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	c.Bind(&l)
	nl, _ := service.StoreLog(*l)
	return c.JSON(http.StatusOK, nl)
}

func UpdateLog(c echo.Context) error {
	l, _ := service.UpdateLog(c)
	return c.JSON(http.StatusOK, l)
}

func GetLogsByItemId(c echo.Context) error {
	ls, _ := service.GetLogsByItemId(c.Param("id"))
	return c.JSON(http.StatusOK, ls)
}

func GetDailyLog(c echo.Context) error {
	ls, err := service.GetDailyLog()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error)
	}
	return c.JSON(http.StatusOK, ls)
}

func GetLogByType(c echo.Context) error {
	ls, err := service.GetLogByType(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error)
	}

	return c.JSON(http.StatusOK, ls)
}
