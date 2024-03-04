package handlers

import (
	"context"

	"github.com/Cahllagerfeld/go-htmx-first-steps/view/pages"
	"github.com/labstack/echo/v4"
)

type AboutHandler struct {
}

func NewAboutHandler() *AboutHandler {
	return &AboutHandler{}
}

func (aboutHandler *AboutHandler) AboutHandler(c echo.Context) error {
	component := pages.About()
	return component.Render(context.Background(), c.Response().Writer)
}
