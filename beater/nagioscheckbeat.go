package beater

import (
	"fmt"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/PhaedrusTheGreek/nagioscheckbeat/check"
	"github.com/PhaedrusTheGreek/nagioscheckbeat/config"
)

type Nagioscheckbeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Nagioscheckbeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Nagioscheckbeat) Run(b *beat.Beat) error {
	logp.Info("nagioscheckbeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()

	for _, checkConfig := range bt.config.Checks {

		checkInstance := check.NagiosCheck{}
		checkInstance.Setup(&checkConfig)
		go checkInstance.Run(func(events []common.MapStr) {
			bt.client.PublishEvents(events)
		})

	}

	for {

		select {
		case <-bt.done:
			return nil
		}
	}

	return nil

}

func (bt *Nagioscheckbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
