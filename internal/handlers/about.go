package handlers

import (
	"context"

	"github.com/Cahllagerfeld/go-htmx-first-steps/view"
	"github.com/labstack/echo/v4"
)

type AboutHandler struct {
}

func NewAboutHandler() *AboutHandler {
	return &AboutHandler{}
}

func (aboutHandler *AboutHandler) AboutHandler(c echo.Context) error {
	component := view.About()
	return component.Render(context.Background(), c.Response().Writer)
}
