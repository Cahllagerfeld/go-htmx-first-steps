package main

import (
	"os"

	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/auth"
	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/handlers"
	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/router"
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

	homeHandler := handlers.NewHomeHandler(e)
	aboutHandler := handlers.NewAboutHandler()
	authHandler := handlers.NewAuthHandler(e)

	auth.NewAuth()

	h := handlers.NewHandler(homeHandler, aboutHandler, authHandler)

	h.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
