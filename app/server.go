package main

import (
	//"fmt"
	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/handler"
	"net/http"
)

func connect(c echo.Context) error {
	database.Connect()
	sqlDb, _ := database.DB.DB()

	defer sqlDb.Close()
	err := sqlDb.Ping()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		return c.String(http.StatusOK, "it works")
	}
}

func main() {
	e := echo.New()

	// Routes
	e.GET("/", connect)
	e.POST("/task/store", handler.CreateType)
	e.GET("/tasks", handler.GetType)
	e.Logger.Fatal(e.Start(":1323"))
}
