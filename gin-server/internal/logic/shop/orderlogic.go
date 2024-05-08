package shop

import (
	"free/gin-server/internal/svc"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderLogic struct {
	logx.Logger
	ctx    *gin.Context
	svcCtx *svc.ServiceContext
}

func NewOrderLogic(ctx *gin.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (o *OrderLogic) Order() {
	//logc.Info(o.ctx, "order logic func")
	//o.Logger.Info("hello i am logic  func !! mother fuck")
	o.ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello order func",
	})
}
