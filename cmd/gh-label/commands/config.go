package commands

import (
	"github.com/erdaltsksn/cui"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var token string

// configCmd represents the config command.
var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "Configure the application",
	Long:    `Set up GitHub token and any configurations needed for this app works.`,
	Example: `gh-label config --token <GITHUB_TOKEN>`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if token == "" {
			cui.Warning(
				"You have to enter a valid GitHub token",
				`Use --token "YOUR_GITHUB_TOKEN" as a flag`,
			)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := keyring.Set("gh-label", "anon", token); err != nil {
			cui.Error("Error while saving the token", err)
		}

		cui.Success("The GitHub Token saved into key manager.")
	},
}

func init() {
	RootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.
	configCmd.Flags().StringVarP(&token, "token", "t", "",
		`A personal access token is required to access private repositories.
You can generate your token here: https://github.com/settings/tokens/new`)
}
