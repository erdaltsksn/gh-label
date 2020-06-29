package main

import (
	"github.com/erdaltsksn/gh-label/cmd"
)

var version = "unknown"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
