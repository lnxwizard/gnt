package cmd

import (
	"github.com/cli/browser"
	"github.com/lnxwizard/gnt/pkg/utils"
	"github.com/spf13/cobra"
)

// suggestion strings for bug command
var suggestForBug = []string{"buug", "bgu", "bg", "bu", "bugg", "issue"}

var bugCmd = &cobra.Command{
	Use:        "bug",
	Short:      "Report bug.",
	Long:       "Report a bug/issue from GitHub.",
	Example:    "gnt bug",
	SuggestFor: suggestForBug,
	Run:        reportBugOnGitHub,
}

func reportBugOnGitHub(cmd *cobra.Command, args []string) {
	// open bug report url in your default browser
	err := browser.OpenURL("https://github.com/lnxwizard/gnt/issues/new")

	// error handling
	utils.HandleError(err)
}

func init() {
	// add `bug` command under root command
	rootCmd.AddCommand(bugCmd)
}
