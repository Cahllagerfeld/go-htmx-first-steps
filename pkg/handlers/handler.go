package handlers

import (
	"context"

	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/view"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Home  *HomeHandler
	About *AboutHandler
}

func NewHandler(home *HomeHandler, about *AboutHandler) *Handler {
	return &Handler{
		Home:  home,
		About: about,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/", h.Home.Index)

	about := e.Group("/about")
	about.GET("", h.About.About)
	about.POST(("/submit"), func(c echo.Context) error {
		component := view.Success()
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.POST("/click", func(c echo.Context) error {
		component := view.Clicked("Tester")
		return component.Render(context.Background(), c.Response().Writer)
	})

}
