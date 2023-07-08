package sv

import (
	"github.com/cli/browser"
	"github.com/lnxwizard/gnt/pkg/utils"
	"github.com/spf13/cobra"
)

func NewCmdSv() *cobra.Command {
	// define sv command
	cmd := &cobra.Command{
		Use:   "sv",
		Short: "Report a security vulnerability",
		Run: func(cmd *cobra.Command, args []string) {
			// reprot security vulnerability
			err := browser.OpenURL("https://github.com/lnxwizard/gnt/security/advisories/new")
			utils.HandleError(err)
		},
	}

	// return cobra command
	return cmd
}
