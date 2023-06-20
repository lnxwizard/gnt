package utils

import (
	"github.com/fatih/color"
	"os"
)

func GetCurrentWorkingDirectory() string {
	// get current working directory
	cwd, err := os.Getwd()
	// error handling
	if err != nil {
		color.Red("Error: Cannot get current working directory.")
		os.Exit(1)
	}

	// return current working directory as type string
	return cwd
}
