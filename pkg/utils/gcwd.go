package utils

import (
	"os"
)

func GetCurrentWorkingDirectory() string {
	// get current working directory
	cwd, err := os.Getwd()
	// error handling
	HandleError(err)

	// return current working directory as type string
	return cwd
}
