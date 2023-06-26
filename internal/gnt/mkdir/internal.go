package mkdir

import (
	"github.com/lnxwizard/gnt/pkg/utils"
	"os"
)

func CreateInternalFolder(projectName string) {
	// create `internal` folder
	err := os.Mkdir(projectName+"/internal", 0755)

	// error handling
	utils.HandleError(err)
}
