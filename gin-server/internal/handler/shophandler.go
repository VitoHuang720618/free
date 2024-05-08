package handler

import (
	"free/gin-server/internal/logic/shop"
	"free/gin-server/internal/svc"
	"github.com/gin-gonic/gin"
)

func shopHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := shop.NewOrderLogic(c, svcCtx)
		l.Order()
		//l := shop.NewOrderLogic(c, svcCtx)
		//l.Order(c)
	}
}
