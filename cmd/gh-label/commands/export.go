package commands

import (
	"strings"

	"github.com/erdaltsksn/cui"
	"github.com/spf13/cobra"

	"github.com/erdaltsksn/gh-label/internal/githubv4"
)

var out string

// exportCmd represents the export command.
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export GitHub label list into a file",
	Long: `Find relevant repository and export label list into a file. You can
state which directory the file will be out using parameter.`,
	Example: `# Export the labels into a file at the current directory
gh-label export --repo erdaltsksn/playground

# Export the labels into a file by specifying absolute file path
gh-label export --repo erdaltsksn/playground --out ~/Desktop/mylabels.json`,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateFlagRepoIsValid()

		if out == "" {
			out = "./" + strings.Replace(repo, "/", "-", -1) + "-labels.json"
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		githubv4.ExportLabels(repo, out)

		cui.Success("The Labels are exported into a file:", out)
	},
}

func init() {
	RootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.
	exportCmd.Flags().StringVarP(&out, "out", "o", "",
		`Output file which contains label list will be save here.
Use 'directory/filename.json' format`)
}
