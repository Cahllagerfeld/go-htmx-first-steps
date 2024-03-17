package domain

import "github.com/shurcooL/githubv4"

type SearchResult struct {
	Search struct {
		IssueCount githubv4.Int
		PageInfo   struct {
			HasNextPage bool
			EndCursor   githubv4.String
		}
		Edges []SearchResultNode
	} `graphql:"search(query: $query, type: ISSUE, first: $pageSize, after: $afterCursor)"`
}

type SearchResultNode struct {
	Node struct {
		PullRequest struct {
			ID     githubv4.ID
			Number githubv4.Int
			Title  githubv4.String
			URL    githubv4.String
			Author struct {
				Login     githubv4.String
				AvatarUrl githubv4.String
			}
			Additions  githubv4.Int
			Deletions  githubv4.Int
			Repository struct {
				NameWithOwner githubv4.String
			}
		} `graphql:"... on PullRequest"`
	} `graphql:"node"`
}
