package main

import (
	"github.com/gabrielhamasaki/htmx-starter/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/public", "public")

	e.GET("/", handlers.Home)

	e.Logger.Fatal(e.Start(":4321"))
}
