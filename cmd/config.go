package cmd

import (
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var token string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure gh-label application",
	Long: `You can use this command to set up GitHub token and any configuration
needed for this app works flawless.`,
	Run: func(cmd *cobra.Command, args []string) {
		if token == "" {
			color.Danger.Println("You have to enter a valid GitHub token")
			color.Info.Prompt(`Use --token "YOUR_GITHUB_TOKEN" as a flag`)
			os.Exit(1)
		}

		if err := keyring.Set("gh-label", "anon", token); err != nil {
			color.Danger.Println("Error while saving the token")
			color.Warn.Prompt(err.Error())
			os.Exit(1)
		}

		color.Success.Prompt("The GitHub Token saved into key manager.")
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
