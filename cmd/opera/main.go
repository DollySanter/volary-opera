package main

import (
	"fmt"
	"os"

	"github.com/DollySanter/volary-opera/cmd/opera/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
