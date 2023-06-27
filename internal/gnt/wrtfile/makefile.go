package wrtfile

import (
	"github.com/lnxwizard/gnt/pkg/utils"
	"os"
)

func WriteMakefile(directory string, projectName string) {
	// write makefile
	err := os.WriteFile(directory, []byte(`run:
	go run cmd/`+projectName+`/main.go

build:
	go build -o bin/`+projectName+` cmd/`+projectName+`/main.go`), 0660)

	// error handling
	utils.HandleError(err)
}
