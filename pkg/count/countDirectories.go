package count

import (
	"github.com/lnxwizard/gnt/pkg/utils"
	"os"
)

func CountDirectories(directory string) int {
	// define number of directories
	numOfDirectories := 1

	// read directory
	directoryEntries, err := os.ReadDir(directory)
	// error handling
	utils.HandleError(err)

	// count directories
	for _, directoryEntry := range directoryEntries {
		if directoryEntry.IsDir() {
			numOfDirectories++
		}
	}

	// return number of directories
	return numOfDirectories
}
