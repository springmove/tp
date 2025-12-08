package echo

import (
	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
)

type Config struct {
	sptty.BaseConfig `yaml:",inline"`

	Port string `yaml:"Port"`
}

func (s *Config) ConfigName() string {
	return base.ServiceEcho
}
