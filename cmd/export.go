package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/erdaltsksn/cui"
	"github.com/erdaltsksn/gh-label/githubv4"
)

var out string
var repo string

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export GitHub label list into a file",
	Long: `Find relevant repository and export label list into a file. You can
state which directory the file will be out using parameter.`,
	Example: `# Export the labels into a file at the current directory
gh-label export --repo erdaltsksn/playground

# Export the labels into a file by specifying absolute file path
gh-label export --repo erdaltsksn/playground --out ~/Desktop/mylabels.json`,
	Run: func(cmd *cobra.Command, args []string) {
		if repo == "" || !strings.Contains(repo, "/") {
			cui.Warning(
				"You have to type the repository name",
				`Use --repo "username/repo-name" as a flag`,
			)
		}

		if out == "" {
			out = "./" + strings.Replace(repo, "/", "-", -1) + "-labels.json"
		}

		githubv4.ExportLabels(repo, out)

		cui.Success("The Labels are exported into a file:", out)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.
	exportCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "",
		`Repository which its labels will be exported into a file.
Please use 'username/repo-name' format.`)
	exportCmd.MarkFlagRequired("repo")
	exportCmd.PersistentFlags().StringVarP(&out, "out", "o", "",
		`Output file which contain label list will be save here.
Use 'directory/filename.json' format`)
}
