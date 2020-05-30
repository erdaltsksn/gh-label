package main

import (
	"log"

	"github.com/spf13/cobra/doc"

	"github.com/erdaltsksn/gh-label/cmd"
)

func main() {
	err := doc.GenMarkdownTree(cmd.GetRootCmd(), "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
