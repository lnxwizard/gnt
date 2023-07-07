package structure

import (
	"github.com/lnxwizard/gnt/pkg/utils"
	"github.com/spf13/cobra"
)

func NewCmdStructure() *cobra.Command {
	// define `structure` command
	cmd := &cobra.Command{
		Use:     "structure",
		Short:   "Print project structure.",
		Long:    "Print structure of your project.",
		Example: "gnt structure",
		RunE: func(cmd *cobra.Command, args []string) error {
			utils.PrintProjectStructure(utils.GetCurrentWorkingDirectory())
			return nil
		},
	}

	// return cobra command
	return cmd
}
