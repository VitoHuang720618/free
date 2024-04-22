package user

import (
	"net/http"

	"micro/gin-server/internal/svc"

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
	g.JSON(http.StatusOK, gin.H{
		"message": "login function !!",
	})
}

func (l *Login) ChangePasswd(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "change password function !!",
	})
}
