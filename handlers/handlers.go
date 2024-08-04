package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(ctx echo.Context, templ templ.Component) error {
	return templ.Render(ctx.Request().Context(), ctx.Response().Writer)
}
