package create

import (
	"github.com/lnxwizard/gnt/internal/gnt/mkdir"
	"github.com/lnxwizard/gnt/internal/gnt/tchfile"
	"github.com/lnxwizard/gnt/internal/gnt/wrtfile"
	"github.com/lnxwizard/gnt/pkg/utils"
	"github.com/spf13/cobra"
)

/*
	Steps to create a Go app:
	1. Create the project folder.
	2. Create the `bin` folder.
	3. Check the command line flags, if used, create new folder by flag. Example: --pkg and --internal flags
	3. Create the directory `cmd/{projectName}`.
	4. Create the main.go file in the `cmd/projectName` directory and write classic `Hello, World!` program in it.
	6. Create go module (go.mod).
*/

// suggestion strings for create command
var suggestForCreate = []string{"craete", "crea", "createe", "creata", "creat"}

func NewCmdCreate() *cobra.Command {
	// define create command
	cmd := &cobra.Command{
		Use:        "create",
		Short:      "Create Go project.",
		Long:       "Create Go project with default template. This is include cmd/{projectName}/main.go directory, go.mod file and bin folder.",
		Example:    "gnt create myGoProject",
		SuggestFor: suggestForCreate,
		Run:        RunCmdCreate,
	}

	// define flags for create command
	cmd.Flags().Bool("pkg", false, "Add `pkg` folder to your project.")
	cmd.Flags().Bool("internal", false, "Add `internal` folder to your project.")
	cmd.Flags().Bool("makefile", false, "Create Makefile.")

	// return cobra command
	return cmd
}

func RunCmdCreate(cmd *cobra.Command, args []string) {
	// define project name
	projectName := args[0]

	// create project folder
	mkdir.CreateProjectFolder(projectName)

	// crate bin folder
	mkdir.CreateBinaryFolder(projectName)

	// define `pkg` flag value
	flagPkgVal, err := cmd.Flags().GetBool("pkg")
	utils.HandleError(err)

	// define `internal` flag value
	flagInternalVal, err := cmd.Flags().GetBool("internal")
	utils.HandleError(err)

	// define `makefile` flag value
	flagMakefileVal, err := cmd.Flags().GetBool("makefile")
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

	if flagMakefileVal {
		wrtfile.WriteMakefile(projectName+"/Makefile", projectName)
	}

	// create cmd/{projectName} directory
	mkdir.CreateCmdDirectory(projectName)

	// write `Hello, World!` program into main.go file
	wrtfile.WriteMainFile(projectName + "/cmd/" + projectName + "/main.go")

	// create go module
	tchfile.CreateGoModule(projectName)

	// Print the project directory tree
	utils.PrintProjectStructure(projectName)
}
