package services

import (
	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
	"github.com/springmove/tp/src/services/echo"
	goredis "github.com/springmove/tp/src/services/go-redis"
	"github.com/springmove/tp/src/services/gorm"
)

var ServiceEcho base.IServiceEcho
var ServiceGoRedis base.IServiceGoRedis
var ServiceGorm base.IServiceGorm

type Services struct {
	sptty.IServices
}

func (s *Services) Services() sptty.Services {
	ServiceEcho = &echo.Service{}
	ServiceGoRedis = &goredis.Service{}
	ServiceGorm = &gorm.Service{}

	return sptty.Services{
		ServiceEcho.(sptty.IService),
		ServiceGoRedis.(sptty.IService),
		ServiceGorm.(sptty.IService),
	}
}

func (s *Services) Configs() sptty.Configs {
	return sptty.Configs{
		&echo.Config{},
		&goredis.Config{},
		&gorm.Config{},
	}
}
