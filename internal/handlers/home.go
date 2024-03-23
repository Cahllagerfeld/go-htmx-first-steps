package handlers

import (
	"context"

	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/auth"
	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/services"
	"github.com/Cahllagerfeld/go-htmx-first-steps/view/pages"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct {
	githubService *services.GithubService
}

func NewIndexHandler(gs *services.GithubService) *IndexHandler {
	return &IndexHandler{
		githubService: gs,
	}
}

func (indexHandler *IndexHandler) indexHandler(c echo.Context) error {
	username := c.Get(auth.Username_Key).(string)
	token := c.Get(auth.GithubToken).(string)

	query, err := indexHandler.githubService.GetPrsToReview(username, token)
	if err != nil {
		return err
	}

	component := pages.IndexPage(username, *query)
	return component.Render(context.Background(), c.Response().Writer)
}

func (IndexHandler *IndexHandler) LoginHandler(c echo.Context) error {
	component := pages.Login()
	return component.Render(context.Background(), c.Response().Writer)
}
