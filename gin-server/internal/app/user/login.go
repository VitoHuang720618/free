package user

import (
	"context"
	"net/http"

	"free/gin-server/internal/svc"

	grpc_server "free/proto/grpc-server"

	"github.com/gin-gonic/gin"
)

type Login struct {
	svc *svc.Services
}

func NewLogin(svc *svc.Services) *Login {
	return &Login{
		svc: svc,
	}
}

func (l *Login) Login(g *gin.Context) {
	// a, e := l.svc.LoginRpc.CheckSession(context.Background(), &grpc_server.LogingRequest{})
	s, err := l.svc.LoginRpc.Login(context.Background(), &grpc_server.LogingRequest{
		Username: "ooo",
		Passwd:   "oooo",
	})
	if err != nil {
		g.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}
	g.JSON(http.StatusOK, gin.H{
		"is login :": s.IsLogin,
	})
}

func (l *Login) ChangePasswd(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "change password function !!",
	})
}
