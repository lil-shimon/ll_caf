package main

import (
	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/handler"
	"github.com/lil-shimon/workManager/middleware"
)

func main() {
	e := echo.New()
	database.Connect()
	sqlDb, _ := database.DB.DB()
	defer sqlDb.Close()

  e.Use(middleware.LogMiddleware)

	// Routes
	e.POST("/category", handler.CreateCategory)
	e.GET("/categories", handler.GetCategories)
	e.GET("/category/:id", handler.ShowCategory)
	e.PUT("/category/:id", handler.UpdateCategory)
	e.DELETE("/category/:id", handler.DeleteCategory)
	e.GET("/tasks", handler.GetTasks)
	e.GET("/task/:id", handler.ShowTask)
	e.PUT("/task/:id", handler.UpdateTask)
	e.DELETE("/task/:id", handler.DeleteTask)
  e.GET("/types", handler.GetTypes)
  e.GET("/type/:id", handler.ShowType)
  e.POST("/type", handler.StoreType)
  e.PUT("/type/:id", handler.UpdateType)
  e.GET("/items", handler.GetItems)
  e.GET("/item/:id", handler.ShowItem)
  e.POST("/item", handler.StoreItem)
  e.PUT("/item/:id", handler.UpdateItem)
  e.GET("/item/type/:id", handler.GetItemsTypeId)
  e.GET("/logs", handler.GetLogs)
  e.GET("/log/:id", handler.GetLog)
  e.POST("/log", handler.StoreLog)
  e.PUT("/log/:id", handler.UpdateLog)
  e.GET("/log/item/:id", handler.GetLogsByItemId)
  e.GET("log/daily", handler.GetDailyLog)

  e.Logger.Fatal(e.Start(":1323"))
}
