package graphqlquery

import "github.com/shurcooL/githubv4"

type PageInfo struct {
	HasNextPage bool
	EndCursor   string
}

type Repository struct {
	Name        string
	Description string
	CreatedAt   githubv4.DateTime
	UpdatedAt   githubv4.DateTime
	URL         githubv4.URI
}

type Repositories struct {
	PageInfo   PageInfo
	TotalCount int32
	Nodes      []Repository
}

type User struct {
	Repositories Repositories `graphql:"repositories(first: $pageSize, after: $after, orderBy: {field: CREATED_AT, direction: DESC})"`
}

type RepositoryQuery struct {
	User User `graphql:"user(login: $username)"`
}
