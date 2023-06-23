package cmd

import (
	"github.com/lnxwizard/gnt/pkg/utils"
	"github.com/spf13/cobra"
)

// define root command
var rootCmd = &cobra.Command{
	Use:     "gnt",
	Short:   "Create Go projects quickly in terminal.",
	Long:    "Create Go projects quickly in terminal. (Long Message)",
	Example: "gnt --version",
	Version: "0.2.0",
}

// Execute root command
func Execute() {
	err := rootCmd.Execute()
	// error handling
	utils.HandleError(err)
}
