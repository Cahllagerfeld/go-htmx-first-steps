package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

type AuthHandler struct {
	e *echo.Echo
}

func NewAuthHandler(e *echo.Echo) *AuthHandler {
	return &AuthHandler{e: e}
}

const providerKey string = "provider"

func (*AuthHandler) AuthCallback(ctx echo.Context) error {
	provider := ctx.Param("provider")
	user, err := gothic.CompleteUserAuth(ctx.Response(), ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), providerKey, provider)))
	if err != nil {
		return err
	}
	ctx.SetCookie(&http.Cookie{
		Name:     "user",
		Value:    user.AvatarURL,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400,
	})
	return ctx.JSON(200, user)
}

func (*AuthHandler) Logout(ctx echo.Context) error {
	gothic.Logout(ctx.Response(), ctx.Request())
	return ctx.Redirect(302, "/")
}

func (*AuthHandler) Login(ctx echo.Context) error {
	provider := ctx.Param("provider")
	if _, err := gothic.CompleteUserAuth(ctx.Response(), ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), providerKey, provider))); err == nil {
		return ctx.Redirect(302, "/")
	} else {
		gothic.BeginAuthHandler(ctx.Response(), ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), providerKey, provider)))
		return nil
	}
}
