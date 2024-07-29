package config

type BeeShop struct {
	Enable bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	Listen string `mapstructure:"listen" json:"listen" yaml:"listen"`
}
