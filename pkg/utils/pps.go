package utils

import (
	"fmt"
	"github.com/lnxwizard/gnt/pkg/count"
	"github.com/plouc/textree"
	"os"
)

func PrintProjectStructure(projectName string) {
	// print project structure title message
	fmt.Println("Your Project Structure")

	// define tree view and directory
	tree, err := textree.TreeFromDir(projectName)
	// error handling
	HandleError(err)

	// render tree view
	tree.Render(os.Stdout, textree.NewRenderOptions())

	// print number of directories and files
	fmt.Println(count.CountDirectories(projectName), "Directories,", count.CountFiles(projectName), "Files\n ")
}
