package config

type configurationFile struct {
	Server struct {
		Cors bool `yaml:"cors"`
	} `yaml:"server"`
}
