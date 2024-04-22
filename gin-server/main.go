package main

import (
	"flag"
	"fmt"

	"micro/gin-server/internal/config"
	"micro/gin-server/internal/handler"
	"micro/gin-server/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "./etc/ginserver.yaml", "the config file")

/**
* 主要實作
* 1 etc/ginserver.yaml , 需和internal/config/config.go 搭配
* 2 router/router.go register gin function
* 3 app/ 是gin router function
* */
func main() {
	flag.Parse()

	var c config.Config

	conf.MustLoad(*configFile, &c)

	svc := svc.NewServices(c)

	r := handler.RegisterHandlers(svc)

	fmt.Printf("Starting server at %s:%s...\n", c.Host, c.Port)

	r.Run(c.Port)
}
