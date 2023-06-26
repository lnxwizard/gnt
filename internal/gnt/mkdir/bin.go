package mkdir

import (
	"github.com/lnxwizard/gnt/pkg/utils"
	"os"
)

func CreateBinaryFolder(projectName string) {
	// create `bin` directory
	err := os.Mkdir(projectName+"/bin", 0755)

	// error handling
	utils.HandleError(err)
}
