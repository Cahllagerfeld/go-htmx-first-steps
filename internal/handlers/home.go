package handlers

import (
	"context"
	"fmt"

	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/auth"
	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/domain"
	"github.com/Cahllagerfeld/go-htmx-first-steps/view/pages"
	"github.com/labstack/echo/v4"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (indexHandler *IndexHandler) indexHandler(c echo.Context) error {
	username := c.Get(auth.Username_Key).(string)
	token := c.Get(auth.GithubToken).(string)
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)

	var query domain.SearchResult

	pageSize := 10
	var afterCursor *githubv4.String

	variables := map[string]interface{}{
		"query":       githubv4.String(fmt.Sprintf("is:open review-requested:%s", username)),
		"pageSize":    githubv4.Int(pageSize),
		"afterCursor": afterCursor,
	}
	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
	}

	for _, pr := range query.Search.Edges {
		fmt.Printf("%s: %s\n", pr.Node.PullRequest.Repository.NameWithOwner, pr.Node.PullRequest.URL)
	}

	component := pages.IndexPage(username, query)
	return component.Render(context.Background(), c.Response().Writer)
}

func (IndexHandler *IndexHandler) LoginHandler(c echo.Context) error {
	component := pages.Login()
	return component.Render(context.Background(), c.Response().Writer)
}
