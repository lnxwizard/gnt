package cmd

import (
	"github.com/lnxwizard/gnt/internal/gnt/mkdir"
	"github.com/lnxwizard/gnt/internal/gnt/wrtfile"
	"github.com/lnxwizard/gnt/pkg/utils"
	"github.com/spf13/cobra"
)

/*
	Steps to create a Go app:
	1. Create the project folder.
	2. Create the `bin` folder.
	3. Create the directory `cmd/projectName`.
	4. Create the main.go file in the `cmd/projectName` directory.
	5. Check the command line flags, if used, create new folder by flag. Example: --pkg and --internal flags
*/

// suggest `create` command for this arrays items
var suggestStrings = []string{"craete", "crea", "createe", "creata", "creat"}

// define `create` command
var createCmd = &cobra.Command{
	Use:        "create",
	Short:      "Create Go project.",
	Long:       "Create Go project with default template. This is include cmd/{projectName}/main.go directory, go.mod file and bin folder.",
	Example:    "gnt create myGoProject",
	SuggestFor: suggestStrings,
	Run:        createRun,
}

func createRun(cmd *cobra.Command, args []string) {
	// define project name
	projectName := args[0]

	// create project folder
	mkdir.CreateProjectFolder(projectName)

	// define `pkg` flag value
	flagPkgVal, err := cmd.Flags().GetBool("pkg")
	utils.HandleError(err)

	// define `internal` flag value
	flagInternalVal, err := cmd.Flags().GetBool("internal")
	utils.HandleError(err)

	// if pkg flag is used create pkg folder else if internal flag is used create internal folder
	if flagPkgVal && !flagInternalVal {
		mkdir.CreatePkgFolder(projectName)
	} else if !flagPkgVal && flagInternalVal {
		mkdir.CreateInternalFolder(projectName)
	} else if flagPkgVal && flagInternalVal {
		mkdir.CreatePkgFolder(projectName)
		mkdir.CreateInternalFolder(projectName)
	}

	// create cmd/{projectName} directory
	mkdir.CreateCmdDirectory(projectName)

	// write `Hello, World!` program into main.go file
	wrtfile.WriteMainFile(projectName + "/cmd/" + projectName + "/main.go")

	// crate bin folder
	mkdir.CreateBinaryFolder(projectName)

	// Print the project directory tree
	utils.PrintProjectStructure(projectName)
}

// init function for create command
func init() {
	// add `create` command under root command
	rootCmd.AddCommand(createCmd)

	// add `pkg` flag under `create` command
	createCmd.Flags().Bool("pkg", false, "Add `pkg` folder to your project.")
	createCmd.Flags().Bool("internal", false, "Add `internal` folder to your project.")
}
