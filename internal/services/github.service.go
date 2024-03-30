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

func (githubService *GithubService) CreateClient(ctx context.Context, token string) *githubv4.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	httpClient := oauth2.NewClient(ctx, src)
	client := githubv4.NewClient(httpClient)
	return client
}

func (githubService *GithubService) GetPrsToReview(client *githubv4.Client, username string, pageParams GithubPaginationParams) (*graphqlquery.ReviewSearchResult, error) {
	var query graphqlquery.ReviewSearchResult
	variables := map[string]interface{}{
		"query":    githubv4.String(fmt.Sprintf("is:open is:pull-request review-requested:%s", username)),
		"pageSize": githubv4.Int(pageParams.PageSize),
	}

	if pageParams.After != "" {
		variables["after"] = githubv4.String(pageParams.After)
	} else {
		variables["after"] = (*githubv4.String)(nil)

	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}
	return &query, nil
}
