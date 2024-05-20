package main

import (
	"flag"
	"fmt"

	"free/gin-server/internal/config"
	"free/gin-server/internal/handler"
	"free/gin-server/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "./etc/ginserver.yaml", "the config file")

func main() {
	flag.Parse()

	// var c config.Config
	var c config.Config

	conf.MustLoad(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)

	r := handler.RegisterHandlers(svcCtx)

	fmt.Printf("Starting server at %s:%s...\n", c.Host, c.Port)

	r.Run(c.Port)
}
