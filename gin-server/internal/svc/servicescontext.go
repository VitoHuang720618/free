package svc

import (
	"free/gin-server/internal/config"
)

type ServiceContext struct {
	Config config.Config
	// LoginRpc grpc_server.LoginClient
	MySQL *MySqlContext
	Redis *RedisContext
}

func NewServiceContext(c config.Config) *ServiceContext {
	MySql := NewMySqlClient(c)
	Redis := NewRedisClient(c)
	return &ServiceContext{
		Config: c,
		// LoginRpc: grpc_server.NewLoginClient(zrpc.MustNewClient(c.Login).Conn()),
		MySQL: MySql,
		Redis: Redis,
	}
}
