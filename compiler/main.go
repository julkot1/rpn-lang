package main

import (
	"fmt"
	"os"
	"rpn/cli"
)

func main() {

	// Parse STC line
	Cli := cli.CreateCli()
	if err := Cli.Root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
