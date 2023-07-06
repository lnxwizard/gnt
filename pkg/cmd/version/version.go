package version

import (
	"github.com/spf13/cobra"
	"go.szostok.io/version/extension"
)

func NewCmdVersion() *cobra.Command {
	cmd := extension.NewVersionCobraCmd(
		extension.WithUpgradeNotice("lnxwizard", "gnt"),
	)

	return cmd
}
