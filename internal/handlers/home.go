package handlers

import (
	"context"
	"fmt"

	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/auth"
	"github.com/Cahllagerfeld/go-htmx-first-steps/view/pages"
	"github.com/google/go-github/v60/github"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (indexHandler *IndexHandler) indexHandler(c echo.Context) error {
	username := c.Get(auth.Username_Key).(string)
	token := c.Get(auth.GithubToken).(string)
	client := github.NewClient(nil).WithAuthToken(token)
	reviewRequests, _, err := client.Search.Issues(context.Background(), fmt.Sprintf("is:open review-requested:%s", username), nil)

	if err != nil {
		c.Logger().Errorf("Error fetching Pull Requests: %v", err)
	}

	for _, pr := range reviewRequests.Issues {
		fmt.Println(pr.GetHTMLURL())
	}

	component := pages.IndexPage(username)
	return component.Render(context.Background(), c.Response().Writer)
}

func (IndexHandler *IndexHandler) LoginHandler(c echo.Context) error {
	component := pages.Login()
	return component.Render(context.Background(), c.Response().Writer)
}
