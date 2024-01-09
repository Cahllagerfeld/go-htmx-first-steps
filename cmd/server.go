package main

import (
	"context"

	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/handlers"
	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/router"
	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/view"
	"github.com/labstack/echo/v4"
)

func main() {
	e := router.New()

	hh := handlers.NewHomeHandler(e)
	ah := handlers.NewAboutHandler()

	h := handlers.NewHandler(hh, ah)

	e.GET("/", h.Home.Index)

	e.GET("/about", h.About.About)

	e.POST(("/about/submit"), func(c echo.Context) error {
		component := view.Success()
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.POST("/click", func(c echo.Context) error {
		component := view.Clicked("Tester")
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
