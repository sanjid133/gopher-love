package main

import (
	"github.com/sanjid133/gopher-love/cmd/gopher"
	"os"
)

func main() {
	//
	if err := gopher.RootCmd().Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
