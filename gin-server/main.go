package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"free/gin-server/internal/config"
	"free/gin-server/internal/handler"
	"free/gin-server/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "./etc/ginserver.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	conf.MustLoad(*configFile, &c)

	// logx
	logx.MustSetup(c.Log)

	// svc context
	svcCtx := svc.NewServiceContext(c)

	// ip address
	address := net.JoinHostPort(svcCtx.Config.Host, fmt.Sprintf("%d", svcCtx.Config.Port))
	if address == "" {
		logx.Error("err listen addr")
	}

	r := handler.NewHandler(svcCtx).RegisterRouter()

	server := &http.Server{
		Addr:    address,
		Handler: r,
	}

	go func() {
		logx.Infof("Starting server at : %s", address)
		if err := server.ListenAndServe(); err != nil {
			logx.Errorf("Listen server err : %v", err)
		}
	}()

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		s := <-osSignal
		logx.Infof("[main]  %s", s.String())
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			if err := server.Shutdown(ctx); err != nil {
				logx.Errorf("【main】 異常退出 err : %s", err.Error())
			} else {
				logx.Info("【main】 正常退出")
			}
			logx.Close()
			cancel()
			return
		case syscall.SIGHUP:
			logx.Info("SIGHUP, ignore")
		default:
			logx.Info("other signal")
		}
	}
}
