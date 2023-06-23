package utils

import (
	"fmt"
	"github.com/plouc/textree"
	"os"
)

func PrintProjectStructure() {
	// define tree view and directory
	tree, err := textree.TreeFromDir(".")
	// error handling
	HandleError(err)

	// render tree view
	tree.Render(os.Stdout, textree.NewRenderOptions())

	// print number of directories and files
	fmt.Println(CountDirectories("."), "Directories,", CountFiles("."), "Files\n ")
}
