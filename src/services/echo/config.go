package echo

import (
	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
)

type Config struct {
	sptty.BaseConfig

	Enable bool   `yaml:"enable"`
	Port   string `yaml:"port"`
}

func (s *Config) ConfigName() string {
	return base.ServiceEcho
}

func (s *Config) Default() sptty.IConfig {
	return &Config{
		Enable: false,
	}
}
