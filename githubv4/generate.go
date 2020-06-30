package githubv4

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/erdaltsksn/cui"
)

// GenerateLabels will generate a label list and put them into a repository.
func GenerateLabels(repository string, fileLabel io.Reader) {
	repo := strings.Split(repository, "/")

	data := graphqlQuery(fmt.Sprintf(`
		query {
			repository(owner: "%s", name: "%s") {
				id
			}
		}
	`, repo[0], repo[1]))

	var labels Nodes
	if err := json.NewDecoder(fileLabel).Decode(&labels); err != nil {
		cui.Error("We couldn't marshal the labels file", err)
	}

	for i := range labels {
		_ = graphqlQuery(fmt.Sprintf(`
			mutation {
				createLabel(input: {
					repositoryId: "%s"
					name: "%s"
					color: "%s"
					description: "%s"
				}) {
					clientMutationId
				}
			}
	`, data.Repository.ID, labels[i].Name, labels[i].Color, labels[i].Description))
	}
}
