// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type NagiosCheckConfig struct {
	Cmd     *string       `yaml:"cmd"`
	Args    *string       `yaml:"args"`
	Name    *string       `yaml:"name"`
	Period  time.Duration `yaml:"period"`
	Enabled *bool         `yaml:"enabled"`
}

type Config struct {
	Checks []NagiosCheckConfig
}


var DefaultConfig = Config{}
