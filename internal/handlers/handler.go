package handlers

import (
	"context"

	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/middleware"
	"github.com/Cahllagerfeld/go-htmx-first-steps/view"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", indexHandler)

	about := e.Group("/about", middleware.WithAuth)
	about.GET("", aboutHandler)
	about.POST(("/submit"), func(c echo.Context) error {
		component := view.Success()
		return component.Render(context.Background(), c.Response().Writer)
	})

	auth := e.Group("/auth")
	auth.GET("/:provider/callback", authCallbackHandler)
	auth.GET("/:provider", loginHandler)
	auth.GET("/logout", logoutHandler)

	e.POST("/click", func(c echo.Context) error {
		component := view.Clicked("Tester")
		return component.Render(context.Background(), c.Response().Writer)
	})

}
