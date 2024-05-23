package shop

import (
	"net/http"

	"free/gin-server/internal/entities/shopentities"
	"free/gin-server/internal/logic/shop"
	"free/gin-server/internal/svc"

	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/breaker"
)

func OrderHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1 http request response struct
		var req shopentities.OrderReq

		// 2 parse http request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 3 call logic
		if err := breaker.DoWithFallback(c.Request.URL.Path, func() error {
			l := shop.NewOrderLogic(c, svcCtx)
			resp, err := l.Order(&req)
			if err != nil {
				return err
			}
			c.JSON(http.StatusOK, gin.H{"message": resp})
			return nil
		}, func(err error) error {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Service is currently unavailable"})
			return nil
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}
