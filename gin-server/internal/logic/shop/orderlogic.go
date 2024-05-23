package shop

import (
	"errors"
	"free/gin-server/internal/entities/shopentities"
	"free/gin-server/internal/svc"
	"math/rand"
	"time"

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
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	num := r.Intn(100)
	if num%4 == 0 {
		return resp, nil
	} else if num%5 == 0 {
		return nil, errors.New("555")
	} else {
		return nil, errors.New("dummy")
	}
	// return resp, nil
}
