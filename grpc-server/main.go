package main

import (
	"flag"
	"fmt"
	"free/grpc-server/internal/server"
	"free/grpc-server/internal/svc"
	grpc_server "free/proto/grpc-server"

	"free/grpc-server/internal/config"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	conf.MustLoad(*configFile, &c)

	svc := svc.NewServices(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		// register grpc server ---> 在server/底下
		grpc_server.RegisterHealthServer(grpcServer, server.NewHealth(svc))
		grpc_server.RegisterLoginServer(grpcServer, server.NewLogin(svc))
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
