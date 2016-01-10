package check

import (
	"github.com/PhaedrusTheGreek/nagioscheckbeat/check"
	"github.com/PhaedrusTheGreek/nagioscheckbeat/config"
	"github.com/PhaedrusTheGreek/nagioscheckbeat/module"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/helper"
)

func init() {
	Metric.Register()
}

// Metric object
var Metric = helper.NewMetric("check", Check{}, module.Module)

var Config = &config.NagiosCheckConfig{}

type Check struct {
	helper.MetricConfig
	Config config.NagiosCheckConfig
}

func (e Check) Setup() {
	Metric.LoadConfig(&Config)
}

func (e Check) Fetch() []common.MapStr {
	check := check.NagiosCheck{}
	check.Setup(Config)
	events, err := check.Check()
	if err != nil {
		logp.Err("Error On Command: %q: %v", Config.Name, err)
	}
	return events
}
