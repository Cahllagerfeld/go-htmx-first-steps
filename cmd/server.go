package main

import (
	"os"

	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/auth"
	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/handlers"
	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/router"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	_ "github.com/joho/godotenv/autoload"
)

var (
	sessionSecret = os.Getenv("SESSION_SECRET")
	cookieStore   = sessions.NewCookieStore([]byte(sessionSecret))
)

func init() {
	cookieStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   false,
	}
}

func main() {
	e := router.New()

	e.Use(session.MiddlewareWithConfig(session.Config{
		Store: cookieStore,
	}))

	auth.NewAuth()

	handlers.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
