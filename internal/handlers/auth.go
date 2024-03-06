package handlers

import (
	"context"
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

const (
	tokenDuration time.Duration = time.Hour * 24 * 7
)

type AuthService interface {
	FindOrCreateUser(user goth.User) *domain.User
}

type AuthHandler struct {
	UserService AuthService
}

func NewAuthHandler(us AuthService) *AuthHandler {
	return &AuthHandler{
		UserService: us,
	}
}

func (ah *AuthHandler) authCallbackHandler(ctx echo.Context) error {
	provider := ctx.Param("provider")
	user, err := gothic.CompleteUserAuth(ctx.Response(), ctx.Request().WithContext(context.WithValue(context.Background(), "provider", provider)))
	if err != nil {
		return err
	}

	u := ah.UserService.FindOrCreateUser(user)

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
		auth.GithubToken:  user.AccessToken,
	}

	sess.Save(ctx.Request(), ctx.Response())

	return ctx.Redirect(http.StatusFound, "/")
}

func (authHandler *AuthHandler) logoutHandler(ctx echo.Context) error {
	sess, _ := session.Get(auth.SessionName, ctx)
	sess.Options = &sessions.Options{
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
	}
	sess.Values[auth.AuthKey] = false

	sess.Save(ctx.Request(), ctx.Response())
	gothic.Logout(ctx.Response(), ctx.Request())
	return ctx.Redirect(http.StatusFound, "/")
}

func (authHandler *AuthHandler) loginHandler(ctx echo.Context) error {
	provider := ctx.Param("provider")
	if _, err := gothic.CompleteUserAuth(ctx.Response(), ctx.Request().WithContext(context.WithValue(context.Background(), "provider", provider))); err == nil {
		return ctx.Redirect(http.StatusFound, "/")
	} else {
		gothic.BeginAuthHandler(ctx.Response(), ctx.Request().WithContext(context.WithValue(context.Background(), "provider", provider)))
		return nil
	}
}
