package handlers

import (
	"context"

	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/auth"
	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/graphqlquery"
	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/services"
	"github.com/Cahllagerfeld/go-htmx-first-steps/view/pages"
	"github.com/labstack/echo/v4"
	"github.com/shurcooL/githubv4"
)

type SetupService interface {
	CreateClient(ctx context.Context, token string) *githubv4.Client
	GetUserRepositories(client *githubv4.Client, username string, pageParams services.GithubPaginationParams) (*graphqlquery.RepositoryQuery, error)
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

	username := c.Get(auth.Username_Key).(string)
	token := c.Get(auth.GithubToken).(string)

	afterCursor := c.QueryParam("after")

	client := setupHander.githubService.CreateClient(context.Background(), token)

	query, err := setupHander.githubService.GetUserRepositories(client, username, services.GithubPaginationParams{PageSize: 10, After: afterCursor})

	if err != nil {
		return err
	}

	component := pages.SetupPage(query.User.Repositories.Nodes, query.User.Repositories.TotalCount, query.User.Repositories.PageInfo.EndCursor, query.User.Repositories.PageInfo.HasNextPage)
	return component.Render(context.Background(), c.Response().Writer)

}
