package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/auth"
	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/domain"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

var users = []*domain.User{}

const (
	providerKey   string        = "provider"
	tokenDuration time.Duration = time.Hour * 24 * 7
)

func authCallbackHandler(ctx echo.Context) error {
	provider := ctx.Param("provider")
	user, err := gothic.CompleteUserAuth(ctx.Response(), ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), providerKey, provider)))
	if err != nil {
		return err
	}

	u := findOrCreateUser(user)

	sess, _ := session.Get(auth.SessionName, ctx)
	sess.Options = &sessions.Options{
		MaxAge:   int(tokenDuration.Seconds()),
		Path:     "/",
		HttpOnly: true,
	}

	sess.Values = map[interface{}]interface{}{
		auth.AuthKey:      true,
		auth.User_Id_Key:  u.ID,
		auth.Username_Key: u.Username,
	}

	sess.Save(ctx.Request(), ctx.Response())

	return ctx.Redirect(http.StatusFound, "/")
}

func logoutHandler(ctx echo.Context) error {
	sess, _ := session.Get(auth.SessionName, ctx)
	sess.Options = &sessions.Options{
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
	}
	sess.Values[auth.AuthKey] = false

	fmt.Println("Logging out")
	sess.Save(ctx.Request(), ctx.Response())
	gothic.Logout(ctx.Response(), ctx.Request())
	return ctx.Redirect(http.StatusFound, "/")
}

func loginHandler(ctx echo.Context) error {
	provider := ctx.Param("provider")
	if _, err := gothic.CompleteUserAuth(ctx.Response(), ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), providerKey, provider))); err == nil {
		return ctx.Redirect(http.StatusFound, "/")
	} else {
		gothic.BeginAuthHandler(ctx.Response(), ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), providerKey, provider)))
		return nil
	}
}

func findOrCreateUser(user goth.User) *domain.User {
	for _, v := range users {
		if v.Email == user.Email {
			return v
		}
	}
	u := &domain.User{
		ID:       len(users) + 1,
		Email:    user.Email,
		Username: user.NickName,
		Avatar:   user.AvatarURL,
	}
	users = append(users, u)
	return u
}
