package githubv4

import (
	"fmt"
	"strings"

	"github.com/erdaltsksn/cui"
)

// RemoveLabels will remove all labels in a repository.
func RemoveLabels(repository string) {
	repo := strings.Split(repository, "/")

	resp := graphqlQuery(fmt.Sprintf(`
		query {
			repository(owner: "%s", name: "%s") {
				labels(first:100) {
					totalCount
					nodes {
						id
					}
				}
			}
		}
	`, repo[0], repo[1]))

	removedLabelCount := resp.Repository.Labels.TotalCount
	for i := range resp.Repository.Labels.Nodes {
		graphqlQuery(fmt.Sprintf(`
			mutation DeleteLabels {
				deleteLabel(input: {id: "%s"}) {
					clientMutationId
				}
			}
		`, resp.Repository.Labels.Nodes[i].ID))
	}

	cui.Info(fmt.Sprintf("%d labels are removed", removedLabelCount))
}
