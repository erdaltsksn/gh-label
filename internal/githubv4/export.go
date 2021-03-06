package githubv4

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/erdaltsksn/cui"
)

// ExportLabels will export labels into a file.
func ExportLabels(repository string, out string) {
	repo := strings.Split(repository, "/")

	resp := graphqlQuery(fmt.Sprintf(`
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

	labels, err := json.MarshalIndent(resp.Repository.Labels.Nodes, "", "    ")
	if err != nil {
		cui.Error("We couldn't marshal the graphql query response", err)
	}

	if err := ioutil.WriteFile(out, labels, 0755); err != nil {
		cui.Error("There is a problem with writing labels into the file", err)
	}
}
