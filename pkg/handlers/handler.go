package handlers

import (
	"context"

	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/view"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Home  *HomeHandler
	About *AboutHandler
	Auth  *AuthHandler
}

func NewHandler(home *HomeHandler, about *AboutHandler, auth *AuthHandler) *Handler {
	return &Handler{
		Home:  home,
		About: about,
		Auth:  auth,
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

	e.GET("/auth/:provider/callback", func(c echo.Context) error {
		h.Auth.AuthCallback(c)
		return nil
	})
	e.GET("/auth/:provider", func(c echo.Context) error {
		h.Auth.Login(c)
		return nil
	})
	e.GET("/logout/:provider", func(context echo.Context) error {
		h.Auth.Logout(context)
		return nil
	})

	e.POST("/click", func(c echo.Context) error {
		component := view.Clicked("Tester")
		return component.Render(context.Background(), c.Response().Writer)
	})

}
