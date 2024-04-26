package shop

import (
	"net/http"

	"free/gin-server/internal/svc"

	"github.com/gin-gonic/gin"
)

type OrderStruct struct {
	svcCtx *svc.ServiceContext
}

func NewOrderStruct(svcCtx *svc.ServiceContext) *OrderStruct {
	return &OrderStruct{
		svcCtx: svcCtx,
	}
}

func (o *OrderStruct) Order(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "Hello order func",
	})
}
