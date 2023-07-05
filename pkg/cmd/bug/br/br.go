package br

import (
	"github.com/cli/browser"
	"github.com/lnxwizard/gnt/pkg/utils"
	"github.com/spf13/cobra"
)

func NewCmdBr() *cobra.Command {
	// define `br` command
	cmd := &cobra.Command{
		Use:   "br",
		Short: "Create a bug report on GitHub.",
		Run:   RunCmdBr,
	}

	// return cobra command
	return cmd
}

func RunCmdBr(cmd *cobra.Command, args []string) {
	// open blank issue
	err := browser.OpenURL("https://github.com/lnxwizard/gnt/issues/new?assignees=&labels=&projects=&template=bug_report.md&title=")
	utils.HandleError(err)
}
