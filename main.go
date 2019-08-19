package main

import (
	"os"

	"github.com/PhaedrusTheGreek/nagioscheckbeat/cmd"

	_ "github.com/PhaedrusTheGreek/nagioscheckbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
