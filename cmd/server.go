package main

import (
	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/handlers"
	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/router"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	e := router.New()

	hh := handlers.NewHomeHandler(e)
	ah := handlers.NewAboutHandler()

	h := handlers.NewHandler(hh, ah)

	h.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
