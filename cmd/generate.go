package cmd

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/erdaltsksn/cui"
	"github.com/erdaltsksn/gh-label/githubv4"
)

var force bool
var list string
var file string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate labels using a list",
	Long:  `Generate labels using predefined label list or a custom label file.`,
	Example: `# Generate the labels using a predefined list
gh-label generate --repo erdaltsksn/playground --list "insane"

# User custom file as a list to generate the labels
gh-label generate --repo erdaltsksn/playground --file my-labels.json

# DANGER: Remove all the labels before generating the labels
gh-label generate --repo erdaltsksn/playground --list "insane" --force`,
	Run: func(cmd *cobra.Command, args []string) {
		if repo == "" || !strings.Contains(repo, "/") {
			cui.Warning(
				"You have to type the repository name",
				`Use --repo "username/repo-name" as a flag`,
			)
		}

		if (file == "") && (list == "") {
			cui.Warning(
				"You have to enter either --file or --list",
				`Use --list "ultimate" as a flag`,
			)
		}

		var fileLabel io.Reader
		if file != "" {
			f, err := os.Open(file)
			if err != nil {
				cui.Error("Error while trying to open the labels file", err)
			}
			fileLabel = f
		} else {
			resp, err := http.Get("https://raw.githubusercontent.com/erdaltsksn/gh-label/master/labels/" + list + ".json")
			if err != nil {
				cui.Warning(
					"We couldn't load the predefined labels.",
					`Use --file "my-labels.json" as a flag`,
				)
			}
			fileLabel = resp.Body
		}

		if force {
			githubv4.RemoveLabels(repo)
		}

		githubv4.GenerateLabels(repo, fileLabel)

		cui.Success("The Labels are imported into the repository")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.
	generateCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "",
		`Repository which its labels will be exported into a file.
Please use 'username/repo-name' format.`)
	generateCmd.MarkFlagRequired("repo")
	generateCmd.PersistentFlags().BoolVarP(&force, "force", "f", false,
		`This will remove all labels before generate the labels.`)
	generateCmd.PersistentFlags().StringVarP(&list, "list", "l", "",
		`Predefined label list name. Use --list "ABC"`)
	generateCmd.PersistentFlags().StringVar(&file, "file", "",
		`Use file as a label list. User --list "file.json"`)
}
