package handler

import (
	"micro/gin-server/internal/app/shop"
	"micro/gin-server/internal/svc"

	"github.com/gin-gonic/gin"
)

type Shop struct {
	svc *svc.Services
}

func NewShopRouter(svc *svc.Services) *Shop {
	return &Shop{
		svc: svc,
	}
}

func (l *Shop) Register(g *gin.Engine) {
	group := g.Group("/shop")
	{
		s := shop.NewOrderStruct(l.svc)
		group.GET("/order", s.Order)
	}
}
