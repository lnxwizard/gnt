package utils

import (
	"os"
)

func CountFiles(directory string) int {
	// define number of files
	numOfFiles := 1

	// read directory
	directoryEntries, err := os.ReadDir(directory)
	// error handling
	HandleError(err)

	// count files
	for _, directoryEntry := range directoryEntries {
		if !directoryEntry.IsDir() {
			numOfFiles++
		}
	}

	// return number of files
	return numOfFiles
}
