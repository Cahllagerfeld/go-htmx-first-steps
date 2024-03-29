package handlers

import (
	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/middleware"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	IndexHandler *IndexHandler
	AuthHandler  *AuthHandler
}

func RegisterRoutes(e *echo.Echo, handlers *Handlers) {
	e.GET("/", handlers.IndexHandler.indexHandler, middleware.WithAuth)
	e.GET("/login", handlers.IndexHandler.LoginHandler)

	auth := e.Group("/auth")
	auth.GET("/:provider/callback", handlers.AuthHandler.authCallbackHandler)
	auth.GET("/:provider", handlers.AuthHandler.loginHandler)
	auth.GET("/logout", handlers.AuthHandler.logoutHandler)

}
