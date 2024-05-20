package handler

import (
	"free/gin-server/internal/handler/shop"
	"free/gin-server/internal/svc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandlers(svcCtx *svc.ServiceContext) *gin.Engine {
	app := gin.Default()

	// #[handler/shop/]
	shopGroup := app.Group("/shop")
	{
		shopGroup.GET("/order", shop.OrderHandler(svcCtx))
	}

	// #[handler/user/]
	userGroup := app.Group("/user")
	{
		//懶得寫
		userGroup.GET("/info", func(c *gin.Context) {

		})
	}

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return app
}
