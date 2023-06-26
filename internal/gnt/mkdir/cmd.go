package mkdir

import (
	"github.com/lnxwizard/gnt/pkg/utils"
	"os"
)

func CreateCmdDirectory(projectName string) {
	// create cmd/{projectName} directory
	err := os.MkdirAll(projectName+"/cmd/"+projectName, 0755)

	// error handling
	utils.HandleError(err)
}
