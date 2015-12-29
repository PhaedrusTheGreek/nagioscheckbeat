package beat

type NagiosCheckConfig struct {
	Cmd  *string `yaml:"cmd"`
	Args *string `yaml:"args"`
	Name *string `yaml:"name"`
}

type NagiosCheckBeatConfig struct {
	Interval *string `yaml:"interval"`
	Checks   []NagiosCheckConfig
}

type ConfigSettings struct {
	Input NagiosCheckBeatConfig
}
