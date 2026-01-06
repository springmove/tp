package jwt

import (
	"time"

	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
)

type Config struct {
	sptty.BaseConfig `yaml:",inline"`

	Expiry time.Duration `yaml:"Expiry"`
}

func (s *Config) ConfigName() string {
	return base.ServiceJwt
}

func (s *Config) Default() sptty.IConfig {
	return &Config{
		Expiry: 24 * time.Hour,
	}
}
