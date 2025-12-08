package gorm

import (
	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
)

type DBConfig struct {
	Type    string `yaml:"Type"`
	Name    string `yaml:"Name"`
	User    string `yaml:"User"`
	Pwd     string `yaml:"Pwd"`
	Host    string `yaml:"Host"`
	Port    int    `yaml:"Port"`
	Timeout int    `yaml:"Timeout"`

	// for mysql
	Charset string `yaml:"Charset"`
}

type Config struct {
	sptty.BaseConfig `yaml:",inline"`

	Configs []DBConfig `yaml:"Configs"`
}

func (c *Config) ConfigName() string {
	return base.ServiceGorm
}

func (c *Config) Validate() error {
	return nil
}
