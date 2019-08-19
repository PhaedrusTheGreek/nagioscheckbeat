package main

import (
	"os"

	"github.com/PhaedrusTheGreek/nagioscheckbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
