package svc

import (
	"free/gin-server/internal/config"
	grpc_server "free/proto/grpc-server"

	"github.com/zeromicro/go-zero/zrpc"
)

type Services struct {
	Config   config.Config
	LoginRpc grpc_server.LoginClient
	MySQL    *MySqlClient
	Redis    *RedisClient
}

func NewServices(c config.Config) *Services {
	MySql := NewMySqlClient(c)
	Redis := NewRedisClient(c)
	return &Services{
		Config:   c,
		LoginRpc: grpc_server.NewLoginClient(zrpc.MustNewClient(c.Login).Conn()),
		MySQL:    MySql,
		Redis:    Redis,
	}
}
