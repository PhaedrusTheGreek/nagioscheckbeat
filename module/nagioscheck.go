package module

import (
	"github.com/elastic/beats/metricbeat/helper"
)

func init() {
	Module.Register()
}

var Module = helper.NewModule("nagioscheck", NagiosCheckModule{})

var Config = NagiosCheckModuleConfig{}

type NagiosCheckModuleConfig struct {
	Metrics map[string]interface{}
}

type NagiosCheckModule struct {
	Name   string
	Config NagiosCheckModuleConfig
}

func (e NagiosCheckModule) Setup() {
	Module.LoadConfig(&Config)
}
