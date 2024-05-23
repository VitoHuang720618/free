package svc

import (
	"free/gin-server/internal/config"

	"github.com/redis/go-redis/v9"
)

type RedisContext struct {
	Client *redis.Client
}

func NewRedisClient(c config.Config) *RedisContext {
	return &RedisContext{
		Client: newRedis(&c),
	}
}

func newRedis(c *config.Config) *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr:     c.Redis.Host,
			Password: c.Redis.Password,
			DB:       c.Redis.Db,
		})
}
