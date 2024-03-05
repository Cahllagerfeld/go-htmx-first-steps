package auth

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

const (
	AuthKey      = "authenticated"
	User_Id_Key  = "user_id"
	Username_Key = "username"
	SessionName  = "session"
	GithubToken  = "github_token"
)

func NewAuth(store *sessions.CookieStore) {

	github_client := os.Getenv("GITHUB_CLIENT_ID")
	github_secret := os.Getenv("GITHUB_CLIENT_SECRET")
	github_callback := os.Getenv("GITHUB_CALLBACK_URL")

	gothic.Store = store

	goth.UseProviders(
		github.New(github_client, github_secret, github_callback),
	)
}

func AuthCallback(context echo.Context) {
	user, err := gothic.CompleteUserAuth(context.Response(), context.Request())
	if err != nil {
		return
	}
	context.JSON(200, user)
}
