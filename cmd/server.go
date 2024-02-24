package main

import (
	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/auth"
	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/handlers"
	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/router"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	e := router.New()

	hh := handlers.NewHomeHandler(e)
	ah := handlers.NewAboutHandler()
	authHandler := handlers.NewAuthHandler(e)

	auth.NewAuth()

	h := handlers.NewHandler(hh, ah, authHandler)

	h.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
