package handlers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

type AuthHandler struct {
	e *echo.Echo
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var users = []*domain.User{}

const (
	providerKey   string        = "provider"
	tokenDuration time.Duration = 24 * time.Hour
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

	u := findOrCreateUser(user)
	tokenString, err := generateToken(u)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Error signing token: "+err.Error())
	}

	setTokenCookie(ctx, tokenString)

	return ctx.JSON(http.StatusOK, user)
}

func (*AuthHandler) Logout(ctx echo.Context) error {
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

func generateToken(u *domain.User) (string, error) {
	claims := Claims{
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenDuration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signKey := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(signKey))
}

func setTokenCookie(ctx echo.Context, tokenString string) {
	ctx.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		Expires:  time.Now().Add(tokenDuration),
		HttpOnly: true,
		Secure:   true,
	})
}
