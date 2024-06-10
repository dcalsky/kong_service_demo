package config

import (
	"flag"
)

var (
	Conf      = new(Config)
	configDir string
)

type Config struct {
	KongSecret string `yaml:"KongSecret"`
}

func init() {
	flag.StringVar(&configDir, "conf-dir", "", "config dir path")
	flag.Parse()
}

func MustInit() {
	err := unmarshalConfDir(configDir, Conf)
	if err != nil {
		panic(err)
	}
}
