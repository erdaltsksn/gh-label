package githubv4

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gookit/color"
)

// ExportLabels will export labels into a file.
func ExportLabels(repository string, out string) {
	repo := strings.Split(repository, "/")

	data := graphqlQuery(fmt.Sprintf(`
		query {
			repository(owner: "%s", name: "%s") {
				id
				labels(first:100) {
					totalCount
					nodes {
						name
						color
						description
					}
				}
			}
		}
	`, repo[0], repo[1]))

	labels, err := json.MarshalIndent(data.Repository.Labels.Nodes, "", "    ")
	if err != nil {
		color.Danger.Println("We couldn't marshal the graphql query response")
		color.Warn.Prompt(err.Error())
		os.Exit(1)
	}

	if err := ioutil.WriteFile(out, labels, 0755); err != nil {
		color.Danger.Println("There is a problem with writing labels into the file")
		color.Warn.Prompt(err.Error())
		os.Exit(1)
	}
}
