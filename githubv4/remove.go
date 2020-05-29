package githubv4

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/machinebox/graphql"
)

// RemoveLabels will remove all labels in a repository.
func RemoveLabels(repository string) {
	repo := strings.Split(repository, "/")

	data := graphqlQuery(fmt.Sprintf(`
		query {
			repository(owner: "%s", name: "%s") {
				labels(first:100) {
					nodes {
						id
					}
				}
			}
		}
	`, repo[0], repo[1]))

	for i := range data.Repository.Labels.Nodes {
		labelGlobalID := data.Repository.Labels.Nodes[i].ID

		graphqlRequest := graphql.NewRequest(fmt.Sprintf(`
			mutation DeleteLabels {
				deleteLabel(input: {id: "%s"}) {
					clientMutationId
				}
			}
		`, labelGlobalID))
		graphqlRequest.Header.Set("Authorization", "bearer "+githubToken)
		graphqlRequest.Header.Set("Accept", "application/vnd.github.bane-preview+json")

		if err := graphqlClient.Run(context.Background(), graphqlRequest, nil); err != nil {
			color.Danger.Println("There is a problem while querying GitHub API v4")
			color.Warn.Prompt(err.Error())
			os.Exit(1)
		}
	}

}
