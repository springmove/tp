package wechat

import (
	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
)

type Service struct {
	sptty.BaseService

	cfg Config

	clients []*wechatClient
}

func (s *Service) ServiceName() string {
	return base.ServiceWechat
}

func (s *Service) Init(app sptty.ISptty) error {
	if err := app.GetConfig(s.ServiceName(), &s.cfg); err != nil {
		return err
	}

	if !s.cfg.Enable {
		sptty.Log(sptty.InfoLevel, "Service Disabled", s.ServiceName())
		return nil
	}

	if err := s.initClients(); err != nil {
		return err
	}

	return nil
}

func (s *Service) initClients() error {

	for k := range s.cfg.Configs {
		client := createWechatClient(&s.cfg.Configs[k])
		if err := client.init(); err != nil {
			return err
		}

		s.clients = append(s.clients, client)
	}

	return nil
}

func (s *Service) Client(index ...int) base.IWechatClient {
	target := 0
	if len(index) > 0 {
		target = index[0]
	}

	return s.clients[target]
}
