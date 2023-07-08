package fr

import (
	"github.com/cli/browser"
	"github.com/lnxwizard/gnt/pkg/utils"
	"github.com/spf13/cobra"
)

func NewCmdFr() *cobra.Command {
	// define `fr` command
	cmd := &cobra.Command{
		Use:   "fr",
		Short: "Report a feature request on GitHub.",
		Run: func(cmd *cobra.Command, args []string) {
			// open blank issue
			err := browser.OpenURL("https://github.com/lnxwizard/gnt/issues/new?assignees=&labels=&projects=&template=feature_request.md&title=")
			utils.HandleError(err)
		},
	}

	// return cobra command
	return cmd
}
