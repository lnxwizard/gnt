package utils

import (
	"fmt"
	"github.com/fatih/color"
)

func HandleError(err error) {
	redFg := color.New(color.FgRed).Add(color.Bold)

	// error handling
	if err != nil {
		redFg.Print("Error:")
		fmt.Println(err)
	}
}
