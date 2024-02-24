package handlers

import (
	"context"

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

func (*AuthHandler) AuthCallback(ctx echo.Context) {
	provider := ctx.Param("provider")
	newRequest := ctx.Request().WithContext(context.WithValue(context.Background(), providerKey, provider))
	user, err := gothic.CompleteUserAuth(ctx.Response(), newRequest)
	if err != nil {
		return
	}
	ctx.JSON(200, user)
}

func (*AuthHandler) Logout(ctx echo.Context) {
	gothic.Logout(ctx.Response(), ctx.Request())
	ctx.Redirect(302, "/")
}

func (*AuthHandler) Login(ctx echo.Context) {
	provider := ctx.Param("provider")
	newRequest := ctx.Request().WithContext(context.WithValue(context.Background(), providerKey, provider))
	gothic.BeginAuthHandler(ctx.Response(), newRequest)
}
