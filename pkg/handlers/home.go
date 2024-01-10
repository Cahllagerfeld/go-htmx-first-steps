package handlers

import (
	"context"

	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/view"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	e *echo.Echo
}

func NewHomeHandler(e *echo.Echo) *HomeHandler {
	return &HomeHandler{
		e: e,
	}
}

func (h *HomeHandler) Index(c echo.Context) error {
	component := view.Hello("Baui")
	h.e.StdLogger.Println("Hello")
	return component.Render(context.Background(), c.Response().Writer)
}
