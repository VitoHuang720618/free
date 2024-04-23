package svc

import (
	"free/gin-server/internal/config"
	// "free/gin-server/internal/svc"

	grpc_server "free/proto/grpc-server"

	"github.com/zeromicro/go-zero/zrpc"
)

type Services struct {
	Config config.Config

	// for grpc client
	// Micro  micro.Service

	// mysql
	Mysql *Mysql
	// grpc
	LoginRpc grpc_server.LoginClient
}

func NewServices(c config.Config) *Services {
	// return nil
	return &Services{
		Config:   c,
		Mysql:    NewMysql(c),
		LoginRpc: grpc_server.NewLoginClient(zrpc.MustNewClient(c.Login).Conn()),
	}
}
