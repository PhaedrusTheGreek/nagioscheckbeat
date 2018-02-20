package beater

import (
	"fmt"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/PhaedrusTheGreek/nagioscheckbeat/check"
	"github.com/PhaedrusTheGreek/nagioscheckbeat/config"
)

type Nagioscheckbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
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

// Beat.Publisher is a called pipeline and is an instance of beat.Pipeline,

func (bt *Nagioscheckbeat) Run(b *beat.Beat) error {
	logp.Info("nagioscheckbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	for _, checkConfig := range bt.config.Checks {

		checkInstance := check.NagiosCheck{}
		checkInstance.Setup(&checkConfig)
		go checkInstance.Run(func(events []beat.Event) {
			bt.client.PublishAll(events)
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
