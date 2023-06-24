package cmd

import (
	"github.com/lnxwizard/gnt/pkg/utils"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// define `pkg` and `internal` flags
var (
	pkgPtr      bool
	internalPtr bool
)

// suggest `create` command for this arrays items
var suggestStrings = []string{"craete", "crea", "createe", "creata", "creat"}

// define `create` command
var createCmd = &cobra.Command{
	Use:        "create",
	Short:      "Create Go project.",
	Long:       "Create Go project with default template. This is include cmd/{projectName}/main.go directory, go.mod file and bin folder.",
	Example:    "gnt create myGoProject",
	SuggestFor: suggestStrings,
	Run: func(cmd *cobra.Command, args []string) {
		// define project name
		projectName := args[0]

		// create base folder
		err := os.Mkdir(projectName, 0755)
		// error handling
		utils.HandleError(err)

		// if flags value is true create `pkg` and `internal` folder
		if pkgPtr == true && internalPtr == false {
			// create directory `pkg`
			err = os.Mkdir(projectName+"/pkg", 0755)
			// error handling
			utils.HandleError(err)
		} else if pkgPtr == false && internalPtr == true {
			// create directory `internal`
			err = os.Mkdir(projectName+"/internal", 0755)
			// error handling
			utils.HandleError(err)
		} else if pkgPtr == true && internalPtr == true {
			// create directory `pkg`
			err = os.Mkdir(projectName+"/pkg", 0755)
			// error handling
			utils.HandleError(err)

			// create directory `internal`
			err = os.Mkdir(projectName+"/internal", 0755)
			// error handling
			utils.HandleError(err)
		}

		// define `cmd` directory name
		cmdDirName := projectName + "/cmd/" + projectName

		// creating directory `cmd/{projectName}`
		err = os.MkdirAll(cmdDirName, 0755)
		// error handling
		utils.HandleError(err)

		// define `main` file directory `cmd/{projectName}/main.go`
		mainFileDir := cmdDirName + "/main.go"

		// create `main.go`
		err = os.WriteFile(mainFileDir, []byte(`package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
`), 0660)
		// error handling
		utils.HandleError(err)

		// create directory `bin`
		err = os.Mkdir(projectName+"/bin", 0755)
		// error handling
		utils.HandleError(err)

		// change current working directory for creating go module file
		err = os.Chdir(utils.GetCurrentWorkingDirectory() + "/" + projectName)

		// define go mod command
		gomodCmd := exec.Command("go", "mod", "init")
		// running go mod command
		err = gomodCmd.Run()

		// error handling
		utils.HandleError(err)

		// Print the directory tree
		utils.PrintProjectStructure(projectName)
	},
}

// init function for create command
func init() {
	// add `create` command under root command
	rootCmd.AddCommand(createCmd)

	// add `pkg` flag under `create` command
	createCmd.Flags().BoolVarP(&pkgPtr, "pkg", "p", false, "Add `pkg` folder to your project.")
	createCmd.Flags().BoolVarP(&internalPtr, "internal", "i", false, "Add `internal` folder to your project.")
}
