package config

type Captcha struct {
	Length int `mapstructure:"length" json:"length" yaml:"length"`
	Expire int `mapstructure:"expire" json:"expire" yaml:"expire"`
	Width  int `mapstructure:"width" json:"width" yaml:"width"`
	Height int `mapstructure:"height" json:"height" yaml:"height"`
}
