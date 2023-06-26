package mkdir

import (
	"github.com/lnxwizard/gnt/pkg/utils"
	"os"
)

func CreateProjectFolder(projectName string) {
	// create project folder
	err := os.Mkdir(projectName, 0755)

	// error handling
	utils.HandleError(err)
}
