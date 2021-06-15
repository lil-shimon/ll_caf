package main

import (
    "github.com/labstack/echo"
    "github.com/lil-shimon/workManager/handler"
)

func main() {
    e := echo.New()

    // Routes
    e.POST("/", handler.CreateType)
    e.GET("/tasks", handler.GetType)
    e.Logger.Fatal(e.Start(":1323"))
}