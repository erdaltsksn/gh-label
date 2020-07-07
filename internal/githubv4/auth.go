package githubv4

import (
	"context"
	"errors"

	"github.com/erdaltsksn/cui"
	"github.com/machinebox/graphql"
	"github.com/zalando/go-keyring"
)

// authCredentials checks if the token is valid or not.
func authCredentials() {
	graphqlRequest := graphql.NewRequest(`
		query {
			viewer {
				login
			}
		}
	`)
	graphqlRequest.Header.Set("Accept", "application/vnd.github.bane-preview+json")

	var resp query
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &resp); err != nil {
		cui.Error("There is a problem while querying GitHub API v4", err)
	}

	if resp.Viewer.Login == "" {
		cui.Error(
			"Bad credentials: You have to enter a valid GitHub token",
			errors.New(`Use --token "YOUR_GITHUB_TOKEN" as a flag`),
		)
	}
}

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
