package irita

import (
	"time"

	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
)

type IritaConfig struct {
	URL    string `yaml:"URL"`
	Key    string `yaml:"Key"`
	Secret string `yaml:"Secret"`

	Timeout time.Duration `yaml:"Timeout"`
}

type Config struct {
	sptty.BaseConfig `yaml:",inline"`

	Configs []IritaConfig `yaml:"Configs"`
}

func (s *Config) ConfigName() string {
	return base.ServiceIrita
}
