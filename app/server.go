package main

import (
	//"fmt"
	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/handler"
	"net/http"
)

func main() {
	e := echo.New()

	database.Connect()
	sqlDb, _ := database.DB.DB()
	defer sqlDb.Close()

	// Routes
	e.POST("/task", handler.CreateType)
	e.GET("/tasks", handler.GetType)
	e.Logger.Fatal(e.Start(":1323"))
}
