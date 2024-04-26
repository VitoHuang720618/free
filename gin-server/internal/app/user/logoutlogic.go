package user

import (
	"context"
	"net/http"

	"free/gin-server/internal/svc"

	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
)

type Logoutlogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogout(ctx context.Context, svcCtx *svc.ServiceContext) *Logoutlogic {
	return &Logoutlogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Logoutlogic) Logout(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "logout function !!",
	})
}
