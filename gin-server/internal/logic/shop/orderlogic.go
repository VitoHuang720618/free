package shop

import (
	"free/gin-server/internal/entities/shopentities"
	"free/gin-server/internal/svc"
	"github.com/gin-gonic/gin"
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

func (o *OrderLogic) Order(req *shopentities.OrderReq) (resp *shopentities.OrderResp, err error) {
	return resp, nil
}
