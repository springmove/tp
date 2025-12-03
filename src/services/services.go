package services

import (
	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
	"github.com/springmove/tp/src/services/echo"
	goredis "github.com/springmove/tp/src/services/go-redis"
	"github.com/springmove/tp/src/services/gorm"
)

type TP struct {
	sptty.IServices

	ServiceEcho    base.IServiceEcho
	ServiceGoRedis base.IServiceGoRedis
	ServiceGorm    base.IServiceGorm
}

func (s *TP) Services() sptty.Services {
	s.ServiceEcho = &echo.Service{}
	s.ServiceGoRedis = &goredis.Service{}
	s.ServiceGorm = &gorm.Service{}

	return sptty.Services{
		s.ServiceEcho.(sptty.IService),
		s.ServiceGoRedis.(sptty.IService),
		s.ServiceGorm.(sptty.IService),
	}
}

func (s *TP) Configs() sptty.Configs {
	return sptty.Configs{
		&echo.Config{},
		&goredis.Config{},
		&gorm.Config{},
	}
}
