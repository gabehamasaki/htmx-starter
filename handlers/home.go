package handlers

import (
	"github.com/gabrielhamasaki/htmx-starter/views"
	"github.com/labstack/echo/v4"
)

func Home(ctx echo.Context) error {

	render(ctx, views.Home())

	return nil
}
