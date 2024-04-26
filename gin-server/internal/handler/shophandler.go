package handler

import (
	"free/gin-server/internal/app/shop"
	"free/gin-server/internal/svc"

	"github.com/gin-gonic/gin"
)

type Shop struct {
	svcCtx *svc.ServiceContext
}

func NewShopRouter(svcCtx *svc.ServiceContext) *Shop {
	return &Shop{
		svcCtx: svcCtx,
	}
}

func (l *Shop) Register(g *gin.Engine) {
	group := g.Group("/shop")
	{
		s := shop.NewOrderStruct(l.svcCtx)
		group.GET("/order", s.Order)
	}
}
