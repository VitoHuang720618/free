package main

import (
	"flag"
	"fmt"
	"free/ws-server/internal/app"
	"free/ws-server/internal/config"
	"free/ws-server/internal/svc"
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "./etc/ws-server.yaml", "the config file")

func main() {
	flag.Parse()

	// var c config.Config
	var c config.Config

	conf.MustLoad(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)

	cm := app.NewConnectionManager(svcCtx)

	//ws server 不會很多，不想搞什麼設計模式
	// app/ 底下接續寫下去就好
	http.HandleFunc("/login", cm.HandleLogin)

	http.HandleFunc("/shop", cm.HandleShop)

	if err := http.ListenAndServe(c.Port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	fmt.Printf("Starting server at %s:%s...\n", c.Host, c.Port)
}
