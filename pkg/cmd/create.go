package cmd

import (
	"github.com/fatih/color"
	"github.com/lnxwizard/gnt/pkg/utils"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// define `create` command
var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create Go project.",
	Long:    "Create Go project with default template. This is include cmd/{projectName}/main.go directory, go.mod file and .out folder.",
	Example: "gnt create myGoProject",
	Run: func(cmd *cobra.Command, args []string) {
		// define project name
		projectName := args[0]

		// create base folder
		err := os.Mkdir(projectName, 0755)
		// error handling
		if err != nil {
			color.Red("Error: Cannot create base directory:", projectName)
			os.Exit(1)
		}

		// define `cmd` directory name
		cmdDirName := projectName + "/cmd/" + projectName

		// creating directory `cmd/{projectName}`
		err = os.MkdirAll(cmdDirName, 0755)
		// error handling
		if err != nil {
			color.Red("Error: Cannot create directory:", cmdDirName)
			os.Exit(1)
		}

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
		if err != nil {
			color.Red("Error: Cannot create directory:", mainFileDir)
			os.Exit(1)
		}

		// create directory `.out`
		err = os.Mkdir(projectName+"/.out", 0755)
		// error handling
		if err != nil {
			color.Red("Error: Cannot create directory: .out")
			os.Exit(1)
		}

		// change current working directory for creating go module file
		err = os.Chdir(utils.GetCurrentWorkingDirectory() + "/" + projectName)

		// create go module file `go.mod`
		gomodCmd := exec.Command("go", "mod", "init")
		// error handling
		if err = gomodCmd.Run(); err != nil {
			color.Red("Error: Cannot run command:", gomodCmd)
			os.Exit(1)
		}
	},
}

// init function for create command
func init() {
	// add `create` command under root command
	rootCmd.AddCommand(createCmd)
}
