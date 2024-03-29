package graphqlquery

import "github.com/shurcooL/githubv4"

type ReviewSearchResult struct {
	Search struct {
		IssueCount githubv4.Int
		PageInfo   struct {
			HasNextPage bool
			EndCursor   string
		}
		Edges []ReviewSearchResultNode
	} `graphql:"search(query: $query, type: ISSUE, first: $pageSize, after: $after)"`
}

type ReviewSearchResultNode struct {
	Node struct {
		PullRequest struct {
			ID     githubv4.ID
			Number int32
			Title  string
			URL    string
			Author struct {
				Login     string
				AvatarUrl string
			}
			Additions  int32
			Deletions  int32
			Repository struct {
				NameWithOwner string
			}
		} `graphql:"... on PullRequest"`
	} `graphql:"node"`
}
