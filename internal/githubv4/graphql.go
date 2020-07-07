package githubv4

import (
	"context"

	"github.com/erdaltsksn/cui"
	"github.com/machinebox/graphql"
	"golang.org/x/oauth2"
)

var graphqlClient = graphql.NewClient("https://api.github.com/graphql",
	graphql.WithHTTPClient(
		oauth2.NewClient(context.Background(),
			oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: getGitHubToken()},
			))),
)

type query struct {
	Viewer struct {
		Login string `json:"login,omitempty"`
	} `json:"viewer,omitempty"`
	Repository struct {
		ID     string `json:"id,omitempty"`
		Labels struct {
			TotalCount int `json:"totalCount,omitempty"`
			Nodes      []struct {
				ID          string `json:"id,omitempty"`
				Name        string `json:"name,omitempty"`
				Color       string `json:"color,omitempty"`
				Description string `json:"description,omitempty"`
			} `json:"nodes,omitempty"`
		} `json:"labels,omitempty"`
	} `json:"repository,omitempty"`
}

func graphqlQuery(q string) query {
	authCredentials()

	graphqlRequest := graphql.NewRequest(q)
	graphqlRequest.Header.Set("Accept", "application/vnd.github.bane-preview+json")

	var resp query
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &resp); err != nil {
		cui.Error("There is a problem while querying GitHub API v4", err)
	}

	return resp
}
