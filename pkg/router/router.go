package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func New() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Static("/assets", "dist")

	return e
}