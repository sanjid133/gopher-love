package main

import (
	"github.com/sanjid133/gopher-love/cmd"
	"os"
)

func main() {
	//
	if err := cmd.RootCmd().Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
