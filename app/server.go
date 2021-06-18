package main

import (
	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/handler"
)

func main() {
	e := echo.New()

	database.Connect()
	sqlDb, _ := database.DB.DB()
	defer sqlDb.Close()

	// Routes
	e.POST("/category", handler.CreateCategory)
	e.GET("/categories", handler.GetCategories)
	e.GET("/category/:id", handler.ShowCategory)
	e.PUT("/category/:id", handler.UpdateCategory)
	e.DELETE("/category/:id", handler.DeleteCategory)
	e.GET("/tasks", handler.GetTasks)
	e.Logger.Fatal(e.Start(":1323"))
}
