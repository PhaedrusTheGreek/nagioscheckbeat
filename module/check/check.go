package check

import (
	"github.com/PhaedrusTheGreek/nagioscheckbeat/beat"
	"github.com/PhaedrusTheGreek/nagioscheckbeat/config"
	"github.com/PhaedrusTheGreek/nagioscheckbeat/module"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/helper"
)

func init() {
	Metric.Register()
}

// Metric object
var Metric = helper.NewMetric("check", Check{}, module.Module)

var Config = &beat.NagiosCheckConfig{}

type Check struct {
	helper.MetricConfig
	Config beat.NagiosCheckConfig
}

func (e Check) Setup() {
	Metric.LoadConfig(&Config)
}

func (e Check) Fetch() []common.MapStr {
	return beat.Check(Config)
}
