package commands

import (
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/erdaltsksn/cui"
	"github.com/erdaltsksn/gh-label/internal/githubv4"
)

var (
	list  string
	file  string
	force bool
)

// generateCmd represents the generate command.
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
	PreRun: func(cmd *cobra.Command, args []string) {
		validateFlagRepoIsValid()

		if (file == "") && (list == "") {
			cui.Warning(
				"You have to enter either --file or --list",
				`Use --list "ultimate" as a flag`,
			)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
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
	generateCmd.Flags().StringVarP(&list, "list", "l", "",
		`Predefined label list name. Use --list "ABC"`)
	generateCmd.Flags().StringVar(&file, "file", "",
		`Use file as a label list. User --list "file.json"`)
	generateCmd.Flags().BoolVarP(&force, "force", "f", false,
		`This will remove all labels before generate the labels.`)
}
