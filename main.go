package main

import (
	nagioscheckbeat "github.com/PhaedrusTheGreek/nagioscheckbeat/beat"
	"github.com/elastic/beats/libbeat/beat"
)

func main() {
	beat.Run("nagioscheckbeat", "0.1", nagioscheckbeat.New())
}
