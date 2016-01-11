package config

type NagiosCheckConfig struct {
	Cmd     *string `yaml:"cmd"`
	Args    *string `yaml:"args"`
	Name    *string `yaml:"name"`
	Period  *string `yaml:"period"`
	Enabled *bool   `yaml:"enabled"`
}

type NagiosCheckBeatConfig struct {
	Checks []NagiosCheckConfig
}

type ConfigSettings struct {
	Input NagiosCheckBeatConfig
}
