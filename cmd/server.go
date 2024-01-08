package main

import (
	"context"

	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/view"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// e.Use(middleware.Logger())

	e.Static("/assets", "dist")

	e.GET("/", func(c echo.Context) error {
		// get headers from context
		headers := c.Request().Header
		e.StdLogger.Println(headers.Get("Accept"))

		component := view.Hello("Baui")
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/about", func(c echo.Context) error {
		component := view.Hello("Tester")
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.POST("/click", func(c echo.Context) error {
		component := view.Clicked("Tester")
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
