package middleware

import (
	"fmt"
	"net/http"

	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/auth"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func WithAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)

		if err != nil {
			fmt.Println("Error getting the session")
		}

		authValue, ok := sess.Values[auth.AuthKey]
		fmt.Println(sess.Values)
		if !ok || authValue != true {
			return c.String(http.StatusUnauthorized, "You are not authorized to view this page")
		}

		return next(c)
	}
}
