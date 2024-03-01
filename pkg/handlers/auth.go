package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/domain"
	"github.com/Cahllagerfeld/go-htmx-first-steps/pkg/auth"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

type AuthHandler struct {
	e *echo.Echo
}

var users = []*domain.User{}

const (
	providerKey   string        = "provider"
	tokenDuration time.Duration = time.Hour * 24 * 7
)

func NewAuthHandler(e *echo.Echo) *AuthHandler {
	return &AuthHandler{e: e}
}

func (*AuthHandler) AuthCallback(ctx echo.Context) error {
	provider := ctx.Param("provider")
	user, err := gothic.CompleteUserAuth(ctx.Response(), ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), providerKey, provider)))
	if err != nil {
		return err
	}

	findOrCreateUser(user)

	sess, _ := session.Get(auth.SessionName, ctx)
	sess.Options = &sessions.Options{
		MaxAge:   int(tokenDuration.Seconds()),
		Path:     "/",
		HttpOnly: true,
	}

	sess.Values[auth.AuthKey] = true
	sess.Save(ctx.Request(), ctx.Response())

	return ctx.Redirect(http.StatusFound, "/")
}

func (*AuthHandler) Logout(ctx echo.Context) error {
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

func (*AuthHandler) Login(ctx echo.Context) error {
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
