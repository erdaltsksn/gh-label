package githubv4

import (
	"context"
	"os"

	"github.com/gookit/color"
	"github.com/machinebox/graphql"
	"github.com/zalando/go-keyring"
)

var graphqlClient = graphql.NewClient("https://api.github.com/graphql")
var githubToken string

func init() {
	token, err := keyring.Get("gh-label", "anon")
	if err != nil {
		color.Danger.Println("Token couldn't load")
		color.Info.Prompt(`Use gh-label config --token "YOUR_GITHUB_TOKEN"`)
		os.Exit(1)
	}

	githubToken = token
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
	graphqlRequest.Header.Set("Authorization", "bearer "+githubToken)
	graphqlRequest.Header.Set("Accept", "application/vnd.github.bane-preview+json")

	var data graphqlQueryResponse
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &data); err != nil {
		color.Danger.Println("There is a problem while querying GitHub API v4")
		color.Warn.Prompt(err.Error())
		os.Exit(1)
	}

	return data
}
