package wrtfile

import (
	"github.com/lnxwizard/gnt/pkg/utils"
	"os"
)

func WriteMainFile(directory string) {
	// write 'Hello, World!' program into main.go
	err := os.WriteFile(directory, []byte(`package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`), 0660)

	// error handling
	utils.HandleError(err)
}
