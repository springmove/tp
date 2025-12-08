package goredis

import (
	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
)

type RedisEntry struct {
	Entry string `yaml:"Entry"`
	Pwd   string `yaml:"Pwd"`

	// only for redis
	DB int `yaml:"DB"`
}

type Config struct {
	sptty.BaseConfig `yaml:",inline"`

	Configs []RedisEntry `yaml:"Configs"`
}

func (s *Config) ConfigName() string {
	return base.ServiceGoRedis
}
