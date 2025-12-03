package goredis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
)

type ClientContext struct {
	base.IClientContext

	ctx    context.Context
	client *redis.Client
}

func (s *ClientContext) Ctx() *context.Context {
	return &s.ctx
}

func (s *ClientContext) Client() *redis.Client {
	return s.client
}

type Service struct {
	sptty.BaseService

	cfg Config

	clients []*ClientContext
}

func (s *Service) ServiceName() string {
	return base.ServiceGoRedis
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
	for _, v := range s.cfg.Configs {

		c := ClientContext{
			ctx: context.Background(),

			client: redis.NewClient(&redis.Options{
				Addr:     v.Entry,
				Password: v.Pwd,
				DB:       v.DB,
			}),
		}

		s.clients = append(s.clients, &c)
	}

	return nil
}

func (s *Service) ClientContext(index ...int) base.IClientContext {
	target := 0
	if len(index) > 0 {
		target = index[0]
	}

	return s.clients[target]
}
