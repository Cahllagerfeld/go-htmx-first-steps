package middleware

import (
	"fmt"
	"net/http"

	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/auth"
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

		if userId, ok := sess.Values[auth.User_Id_Key].(int); ok && userId != 0 {
			c.Set(auth.User_Id_Key, userId)
		}

		if username, ok := sess.Values[auth.Username_Key].(string); ok && len(username) != 0 {
			c.Set(auth.Username_Key, username)
		}

		return next(c)
	}
}
