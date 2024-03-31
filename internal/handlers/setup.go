package handlers

import (
	"context"

	"github.com/Cahllagerfeld/go-htmx-first-steps/view/pages"
	"github.com/labstack/echo/v4"
	"github.com/shurcooL/githubv4"
)

type SetupService interface {
	CreateClient(ctx context.Context, token string) *githubv4.Client
}

type SetupHandler struct {
	githubService SetupService
}

func NewSetupHandler(gs SetupService) *SetupHandler {
	return &SetupHandler{
		githubService: gs,
	}
}

func (setupHander *SetupHandler) GetSetupPage(c echo.Context) error {
	component := pages.SetupPage()
	return component.Render(context.Background(), c.Response().Writer)
}
