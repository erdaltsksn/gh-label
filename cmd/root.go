package cmd

import (
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh-label",
	Short: "This tool helps you generate labels for your Github repositories.",
	Long: `gh-label help you export, import, generate, regenerate or revert the
GitHub Issue Labels to the default ones. There are a few pre-defined label list.
You are welcome to add yours.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Danger.Prompt(err.Error())
		os.Exit(1)
	}
}
