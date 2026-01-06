package wechat

import (
	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
)

type WechatConfig struct {
	sptty.BaseConfig

	Type   string `yaml:"Type"`
	AppID  string `yaml:"AppID"`
	Secret string `yaml:"Secret"`
}

type Config struct {
	sptty.BaseConfig `yaml:",inline"`

	Configs []WechatConfig `yaml:"Configs"`
}

func (s *Config) ConfigName() string {
	return base.ServiceWechat
}
