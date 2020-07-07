package commands

import (
	"strings"

	"github.com/erdaltsksn/cui"
	"github.com/spf13/cobra"
)

// Used by `exportCmd`and `generateCmd`.
var repo string

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "gh-label",
	Short: "This app helps you manage GitHub issue labels.",
	Long: `gh-label helps you export, generate, import, regenerate GitHub Issue
Labels or revert them to the default ones. There are a few pre-defined label list.
You are welcome to add yours.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cui.Error("Something went wrong", err)
	}
}

// GetRootCmd returns the instance of root command.
func GetRootCmd() *cobra.Command {
	return rootCmd
}

func init() {
	rootCmd.AddCommand(cui.VersionCmd)

	rootCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "",
		`Repository which its labels will be generated or exported into a file.
Please use 'username/repo-name' format.`)
}

func validateFlagRepoIsValid() {
	if repo == "" || !strings.Contains(repo, "/") {
		cui.Warning(
			"You have to type the repository name",
			`Use --repo "username/repo-name" as a flag`,
		)
	}
}
