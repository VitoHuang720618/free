package shop

import (
	"net/http"

	"micro/gin-server/internal/svc"

	"github.com/gin-gonic/gin"
)

type OrderStruct struct {
	ctx *svc.Services
}

func NewOrderStruct(svc *svc.Services) *OrderStruct {
	return &OrderStruct{
		ctx: svc,
	}
}

func (o *OrderStruct) Order(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "Hello order func",
	})
}
