package main

import (
	"github.com/sanjid133/gopher-love/cmd/gopher/cmds"
	"os"
)

func main() {
	//
	if err := cmds.RootCmd().Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
