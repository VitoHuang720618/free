package user

import (
	"net/http"

	"micro/gin-server/internal/svc"

	"github.com/gin-gonic/gin"
)

type Logout struct {
	svc *svc.Services
}

func NewLogout(svc *svc.Services) *Logout {
	return &Logout{
		svc: svc,
	}
}

func (l *Logout) Logout(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "logout function !!",
	})
}
