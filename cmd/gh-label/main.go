package main

import (
	"github.com/erdaltsksn/cui"

	"github.com/erdaltsksn/gh-label/cmd/gh-label/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		cui.Error("Something went wrong", err)
	}
}
