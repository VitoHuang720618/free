package shop

import (
	"context"
	"net/http"

	"free/gin-server/internal/svc"

	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (o *OrderLogic) Order(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "Hello order func",
	})
}
