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

type GithubService interface {
	CreateClient(ctx context.Context, token string) *githubv4.Client
	GetPrsToReview(client *githubv4.Client, username string, pageParams services.GithubPaginationParams) (*graphqlquery.ReviewSearchResult, error)
}

type IndexHandler struct {
	githubService GithubService
}

func NewIndexHandler(gs GithubService) *IndexHandler {
	return &IndexHandler{
		githubService: gs,
	}
}

func (indexHandler *IndexHandler) indexHandler(c echo.Context) error {
	username := c.Get(auth.Username_Key).(string)
	token := c.Get(auth.GithubToken).(string)

	client := indexHandler.githubService.CreateClient(context.Background(), token)

	query, err := indexHandler.githubService.GetPrsToReview(client, username, services.GithubPaginationParams{PageSize: 10, After: ""})
	if err != nil {
		return err
	}

	component := pages.IndexPage(username, *query, query.Search.IssueCount)
	return component.Render(context.Background(), c.Response().Writer)
}

func (IndexHandler *IndexHandler) LoginHandler(c echo.Context) error {
	component := pages.Login()
	return component.Render(context.Background(), c.Response().Writer)
}
