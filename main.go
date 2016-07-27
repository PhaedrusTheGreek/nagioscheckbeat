package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/PhaedrusTheGreek/nagioscheckbeat/beater"
)

func main() {
	err := beat.Run("nagioscheckbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
