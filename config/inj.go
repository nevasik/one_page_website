package config

import (
	lg "gitlab.com/nevasik7/logger"
)

var I Inj

type inj struct {
}

type Inj interface {
}

func (cfg *Config) Init() *Config {
	var (
		i = &inj{}
	)

	lg.Infof("Logger connected successfully")

	I = i

	return cfg
}
