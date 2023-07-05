package blank

import (
	"github.com/cli/browser"
	"github.com/lnxwizard/gnt/pkg/utils"
	"github.com/spf13/cobra"
)

func NewCmdBlank() *cobra.Command {
	// define `blank` command
	cmd := &cobra.Command{
		Use:   "blank",
		Short: "Open a blank issue on GitHub.",
		Run:   RunCmdBlank,
	}

	// return cobra command
	return cmd
}

func RunCmdBlank(cmd *cobra.Command, args []string) {
	// open blank issue
	err := browser.OpenURL("https://github.com/lnxwizard/gnt/issues/new")
	utils.HandleError(err)
}
