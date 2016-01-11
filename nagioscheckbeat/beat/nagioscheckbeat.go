package beat

import (
	"github.com/PhaedrusTheGreek/nagioscheckbeat/check"
	"github.com/PhaedrusTheGreek/nagioscheckbeat/config"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

type NagiosCheckBeat struct {
	// Config
	checks []config.NagiosCheckConfig

	// State
	done chan struct{}
}

func New() *NagiosCheckBeat {
	return &NagiosCheckBeat{}
}

func (nagiosCheckBeat *NagiosCheckBeat) Config(b *beat.Beat) error {
	var config config.ConfigSettings
	err := cfgfile.Read(&config, "")
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return err
	}

	nagiosCheckBeat.checks = config.Input.Checks

	return nil
}

func (nagiosCheckBeat *NagiosCheckBeat) Setup(b *beat.Beat) error {

	nagiosCheckBeat.done = make(chan struct{})
	return nil

}

func (nagiosCheckBeat *NagiosCheckBeat) Run(b *beat.Beat) error {

	for _, checkConfig := range nagiosCheckBeat.checks {

		checkInstance := check.NagiosCheck{}
		checkInstance.Setup(&checkConfig)
		go checkInstance.Run(func(events []common.MapStr) {
			b.Events.PublishEvents(events)
		})

	}

	for {

		select {
		case <-nagiosCheckBeat.done:
			return nil
		}
	}

	return nil
}

func (nagiosCheckBeat *NagiosCheckBeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (nagiosCheckBeat *NagiosCheckBeat) Stop() {
	close(nagiosCheckBeat.done)
}
