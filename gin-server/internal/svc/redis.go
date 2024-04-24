package svc

import (
	"free/gin-server/internal/config"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	c    config.Config
	test *redis.Client
}

func NewRedisClient(c config.Config) *RedisClient {
	return &RedisClient{
		test: newRedis(c),
	}
}

func newRedis(c config.Config) *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr:     c.Redis.Host,
			Password: c.Redis.Password,
			DB:       c.Redis.Db,
		})
}
