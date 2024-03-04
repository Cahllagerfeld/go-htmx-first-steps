package handlers

import (
	"context"

	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/middleware"
	"github.com/Cahllagerfeld/go-htmx-first-steps/view/pages"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	AboutHandler *AboutHandler
	IndexHandler *IndexHandler
	AuthHandler  *AuthHandler
}

func RegisterRoutes(e *echo.Echo, handlers *Handlers) {
	e.GET("/", handlers.IndexHandler.indexHandler, middleware.WithAuth)
	e.GET("/login", handlers.IndexHandler.LoginHandler)

	about := e.Group("/about", middleware.WithAuth)
	about.GET("", handlers.AboutHandler.AboutHandler)
	about.POST(("/submit"), func(c echo.Context) error {
		component := pages.Success()
		return component.Render(context.Background(), c.Response().Writer)
	})

	auth := e.Group("/auth")
	auth.GET("/:provider/callback", handlers.AuthHandler.authCallbackHandler)
	auth.GET("/:provider", handlers.AuthHandler.loginHandler)
	auth.GET("/logout", handlers.AuthHandler.logoutHandler)

}
