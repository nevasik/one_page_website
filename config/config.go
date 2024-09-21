package config

import (
	lg "gitlab.com/nevasik7/logger"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	BaseURL string `yaml:"baseURL"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
}

func MustLoad() *Config {
	var (
		cfg Config
	)

	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		lg.Panic(err.Error())
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		lg.Panic(err.Error())
	}

	lg.Infof("Success must load config")

	return &cfg
}
