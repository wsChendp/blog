package config

type Captcha struct {
	Height int `yaml:"height" json:"height"`
	Width  int `yaml:"width" json:"width"`
	Length int `yaml:"length" json:"length"`
	MaxSkew float64 `yaml:"max_skew" json:"max_skew"`
	DotCount int `yaml:"dot_count" json:"dot_count"`
}

