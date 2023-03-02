package main

import (
	"echogorm/interface/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.HelloHandler())
	e.GET("/users", handler.IndexHandler())
	e.GET("/user/:id", handler.ShowHandler())
	e.POST("/users", handler.StoreHandler())
	e.PUT("/user/:id", handler.UpdateHandler())
	e.DELETE("/user/:id", handler.DeleteHandler())
	e.Logger.Fatal(e.Start(":8080"))
}
