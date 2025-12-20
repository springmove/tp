package base

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const (
	ServiceGoRedis = "go-redis"
)

type IClientContext interface {
	Ctx() *context.Context
	Client() *redis.Client
}

var IGoRedis IServiceGoRedis

type IServiceGoRedis interface {
	ClientContext(index ...int) IClientContext
}
