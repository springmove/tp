package irita

import (
	"time"

	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"

	iritaV3 "github.com/bianjieai/avata-sdk-go/v3"
	"github.com/bianjieai/avata-sdk-go/v3/configs"
)

type Service struct {
	sptty.BaseService

	cfg     Config
	clients []*iritaV3.AvataClient
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

func (s *Service) ServiceName() string {
	return base.ServiceIrita
}

func (s *Service) Clients(index ...int) *iritaV3.AvataClient {
	target := 0
	if len(index) > 0 {
		target = index[0]
	}

	return s.clients[target]
}

func (s *Service) initClients() error {
	for _, v := range s.cfg.Configs {

		client := iritaV3.NewClient(
			v.URL,
			v.Key,
			v.Secret,
			[]configs.Options{
				configs.HttpTimeout(8 * time.Second),
			}...,
		)

		s.clients = append(s.clients, client)
	}

	return nil
}
