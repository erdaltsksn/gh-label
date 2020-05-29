package githubv4

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
)

// GenerateLabels will generate a label list and put them into a repository.
func GenerateLabels(repository string, labelsFile string) {
	repo := strings.Split(repository, "/")

	data := graphqlQuery(fmt.Sprintf(`
		query {
			repository(owner: "%s", name: "%s") {
				id
			}
		}
	`, repo[0], repo[1]))

	file, err := os.Open(labelsFile)
	if err != nil {
		color.Danger.Println("Error while trying to open the labels file")
		color.Warn.Prompt(err.Error())
		os.Exit(1)
	}

	var labels Nodes
	if err = json.NewDecoder(file).Decode(&labels); err != nil {
		color.Danger.Println("We couldn't marshal the labels file")
		color.Warn.Prompt(err.Error())
		os.Exit(1)
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
