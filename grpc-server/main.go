package main

import (
	"flag"
	"fmt"
	"free/grpc-server/internal/app"
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
		// register grpc server ---> 在app/底下
		// proto 在proto/grpc-server 底下
		//
		// create proto 語法 + protoc --go_out=./proto --go-grpc_out=./proto ./grpc-server/grpc-server.proto
		grpc_server.RegisterHealthServer(grpcServer, app.NewPing(svc))
		grpc_server.RegisterLoginServer(grpcServer, app.NewLogin(svc))
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
