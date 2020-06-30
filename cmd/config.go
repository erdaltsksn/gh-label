package cmd

import (
	"github.com/erdaltsksn/cui"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var token string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "Configure gh-label application",
	Long:    `Set up GitHub token and any configuration needed for this app works.`,
	Example: `gh-label config --token GITHUB_TOKEN`,
	Run: func(cmd *cobra.Command, args []string) {
		if token == "" {
			cui.Warning(
				"You have to enter a valid GitHub token",
				`Use --token "YOUR_GITHUB_TOKEN" as a flag`,
			)
		}

		if err := keyring.Set("gh-label", "anon", token); err != nil {
			cui.Error("Error while saving the token", err)
		}

		cui.Success("The GitHub Token saved into key manager.")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.
	configCmd.PersistentFlags().StringVarP(&token, "token", "t", "",
		`A personal access token is required to access private repositories.
You can generate your token here: https://github.com/settings/tokens/new`)
	configCmd.MarkFlagRequired("token")
}
