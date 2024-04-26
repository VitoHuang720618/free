package user

import (
	"net/http"

	"free/gin-server/internal/svc"

	"github.com/gin-gonic/gin"
)

type Logout struct {
	svcCtx *svc.ServiceContext
}

func NewLogout(svcCtx *svc.ServiceContext) *Logout {
	return &Logout{
		svcCtx: svcCtx,
	}
}

func (l *Logout) Logout(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "logout function !!",
	})
}
