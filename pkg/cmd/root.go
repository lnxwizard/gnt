package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// define root command
var rootCmd = &cobra.Command{
	Use:     "gnt",
	Short:   "Create Go projects quickly in terminal.",
	Long:    "Create Go projects quickly in terminal. (Long Message)",
	Example: "gnt --version",
	Version: "0.1.0",
}

// Execute root command
func Execute() {
	// error handling
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
