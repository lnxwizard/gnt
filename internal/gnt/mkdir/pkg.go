package mkdir

import (
	"github.com/lnxwizard/gnt/pkg/utils"
	"os"
)

func CreatePkgFolder(projectName string) {
	// create `pkg` folder
	err := os.Mkdir(projectName+"/pkg", 0755)

	// error handling
	utils.HandleError(err)
}
