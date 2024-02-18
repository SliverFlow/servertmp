package config

type System struct {
	Port    int `mapstructure:"port" json:"port" yaml:"port"`
	Timeout int `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
}
