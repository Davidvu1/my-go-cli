package main

import (
	"fmt"
	"os"

	"github.com/Davidvu1/my-go-cli/cmd"
)

func main() {

	err := cmd.Execute()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
