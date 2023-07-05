package root

import (
	"github.com/MakeNowJust/heredoc"
	bugCmd "github.com/lnxwizard/gnt/pkg/cmd/bug"
	createCmd "github.com/lnxwizard/gnt/pkg/cmd/create"
	repoCmd "github.com/lnxwizard/gnt/pkg/cmd/repo"
	"github.com/spf13/cobra"
	"go.szostok.io/version/extension"
)

func NewCmdRoot() *cobra.Command {
	// define root command
	cmd := &cobra.Command{
		Use:   "gnt",
		Short: "Create Go projects quickly in terminal.",
		Long:  "Create Go projects quickly in terminal. (Long Message)",
		Example: heredoc.Doc(`
		$ gnt create myGoApp
		$ gnt bug
		$ gnt repo
		$ gnt create myGoApp --pkg --makefile`),
		Version: "0.6.0-beta",
	}

	// define commands
	cmd.AddCommand(bugCmd.NewCmdBug())
	cmd.AddCommand(createCmd.NewCmdCreate())
	cmd.AddCommand(repoCmd.NewCmdRepo())
	cmd.AddCommand(
		extension.NewVersionCobraCmd(
			extension.WithUpgradeNotice("lnxwizard", "gnt"),
		),
	)

	// return cobra command
	return cmd
}