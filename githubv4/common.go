package githubv4

import (
	"context"
	"errors"

	"github.com/erdaltsksn/cui"
	"github.com/machinebox/graphql"
	"github.com/zalando/go-keyring"
)

var graphqlClient = graphql.NewClient("https://api.github.com/graphql")

func getGitHubToken() string {
	token, err := keyring.Get("gh-label", "anon")
	if err != nil {
		cui.Error("Token couldn't load",
			err,
			errors.New(`Use gh-label config --token "YOUR_GITHUB_TOKEN"`),
		)
	}

	return token
}

type graphqlQueryResponse struct {
	Repository struct {
		ID     string `json:"id,omitempty"`
		Labels struct {
			TotalCount int   `json:"totalCount,omitempty"`
			Nodes      Nodes `json:"nodes,omitempty"`
		} `json:"labels,omitempty"`
	} `json:"repository,omitempty"`
}

// Nodes is a struct for label list
type Nodes []struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Color       string `json:"color,omitempty"`
	Description string `json:"description,omitempty"`
}

func graphqlQuery(q string) graphqlQueryResponse {
	graphqlRequest := graphql.NewRequest(q)
	graphqlRequest.Header.Set("Authorization", "bearer "+getGitHubToken())
	graphqlRequest.Header.Set("Accept", "application/vnd.github.bane-preview+json")

	var data graphqlQueryResponse
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &data); err != nil {
		cui.Error("There is a problem while querying GitHub API v4", err)
	}

	return data
}
