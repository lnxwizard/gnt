package bug

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/cli/browser"
	blankCmd "github.com/lnxwizard/gnt/pkg/cmd/bug/blank"
	brCmd "github.com/lnxwizard/gnt/pkg/cmd/bug/br"
	frCmd "github.com/lnxwizard/gnt/pkg/cmd/bug/fr"
	svCmd "github.com/lnxwizard/gnt/pkg/cmd/bug/sv"
	"github.com/lnxwizard/gnt/pkg/utils"
	"github.com/spf13/cobra"
)

// suggestion strings for bug command
var suggestForBug = []string{"buug", "bgu", "bg", "bu", "bugg", "issue"}

func NewCmdBug() *cobra.Command {
	// define `bug` command
	cmd := &cobra.Command{
		Use:   "bug",
		Short: "Report bug.",
		Long:  "Report a bug/issue from GitHub.",
		Example: heredoc.Doc(`
		$ gnt bug blank
		$ gnt bug br
		$ gnt bug fr`),
		SuggestFor: suggestForBug,
		Run: func(cmd *cobra.Command, args []string) {
			// open bug report url in your default browser
			err := browser.OpenURL("https://github.com/lnxwizard/gnt/issues/new/choose")
			utils.HandleError(err)
		},
	}

	// define subcommands for bug command
	cmd.AddCommand(blankCmd.NewCmdBlank())
	cmd.AddCommand(brCmd.NewCmdBr())
	cmd.AddCommand(frCmd.NewCmdFr())
	cmd.AddCommand(svCmd.NewCmdSv())

	// return cobra command
	return cmd
}
