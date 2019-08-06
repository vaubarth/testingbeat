package main

import (
	"os"

	"github.com/vaubarth/testingbeat/cmd"

	_ "github.com/vaubarth/testingbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
