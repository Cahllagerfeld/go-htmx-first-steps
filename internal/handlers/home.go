package handlers

import (
	"context"

	"github.com/Cahllagerfeld/go-htmx-first-steps/view"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (indexHandler *IndexHandler) indexHandler(c echo.Context) error {
	component := view.Hello("Baui")
	return component.Render(context.Background(), c.Response().Writer)
}
