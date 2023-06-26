package cmd

import (
	"github.com/cli/browser"
	"github.com/lnxwizard/gnt/pkg/utils"
	"github.com/spf13/cobra"
)

// suggestion strings for repo command
var suggestForRepo = []string{"rep", "repository", "source", "reposit", "reposito", "repositor"}

var repoCmd = &cobra.Command{
	Use:        "repo",
	Short:      "Open repository of gnt.",
	Long:       "Open GitHub repository of gnt in your default browser.",
	Example:    "gnt repo",
	SuggestFor: suggestForRepo,
	Run:        openRepository,
}

func openRepository(cmd *cobra.Command, args []string) {
	// open GitHub repository in your default browser
	err := browser.OpenURL("https://github.com/lnxwizard/gnt")

	// error handling
	utils.HandleError(err)
}

func init() {
	// add `repo` command under root command
	rootCmd.AddCommand(repoCmd)
}
