package shop

import (
	"free/gin-server/internal/entities/shopentities"
	"free/gin-server/internal/logic/shop"
	"free/gin-server/internal/svc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OrderHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1
		// entities/shopentities/order.go
		// http request response struct 放這裡
		var req shopentities.OrderReq

		// 2 parse http request
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 3 呼叫logic 業務邏輯程式
		l := shop.NewOrderLogic(c, svcCtx)
		resp, err := l.Order(&req)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": resp,
			})
		}
	}
}
