package services

import (
	"context"
	"fmt"

	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/graphqlquery"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type GithubService struct{}

func NewGithubService() *GithubService {
	return &GithubService{}
}

func (githubService *GithubService) GetPrsToReview(username, token string, pageSize int) (*graphqlquery.ReviewSearchResult, error) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)

	var query graphqlquery.ReviewSearchResult

	var after *githubv4.String

	variables := map[string]interface{}{
		"query":    githubv4.String(fmt.Sprintf("is:open review-requested:%s", username)),
		"pageSize": githubv4.Int(pageSize),
		"after":    after,
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}
	return &query, nil
}
