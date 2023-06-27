package tchfile

import (
	"github.com/lnxwizard/gnt/pkg/utils"
	"os"
	"os/exec"
)

func CreateGoModule(projectName string) {
	// change directory to project name
	err := os.Chdir(utils.GetCurrentWorkingDirectory() + "/" + projectName)
	// error handling
	utils.HandleError(err)

	// create go module
	gomodCmd := exec.Command("go", "mod", "init")

	err = gomodCmd.Run()

	// error handling
	utils.HandleError(err)

	// change directory to back
	err = os.Chdir("..")
	// error handling
	utils.HandleError(err)
}
