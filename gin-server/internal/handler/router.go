package handler

import (
	"free/gin-server/internal/svc"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(svcCtx *svc.ServiceContext) *gin.Engine {
	app := gin.Default()
	shop := app.Group("/shop")
	{
		shop.GET("/order", shopHandler(svcCtx))
	}
	return app
}
