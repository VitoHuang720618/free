package handler

import (
	"fmt"
	"log"
	"net/http"

	"free/gin-server/internal/handler/shop"
	"free/gin-server/internal/svc"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	mredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

type Handler struct {
	svcCtx  *svc.ServiceContext
	limiter *limiter.Limiter
}

func NewHandler(svcCtx *svc.ServiceContext) *Handler {
	rate, err := limiter.NewRateFromFormatted("5-M")
	if err != nil {
		log.Fatal(err)
	}

	var store limiter.Store

	if store, err = mredis.NewStoreWithOptions(svcCtx.Redis.Client, limiter.StoreOptions{
		Prefix:   "limit_example",
		MaxRetry: 3,
	}); err != nil {
		log.Fatal(err)
	}

	return &Handler{
		svcCtx:  svcCtx,
		limiter: limiter.New(store, rate),
	}
}

func MyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Before route")
		ctx.Next()
		fmt.Println("After route")
	}
}

func (r *Handler) RegisterRouter() *gin.Engine {
	app := gin.New()
	h := app.Group("/health")
	{

		h.Use(mgin.NewMiddleware(r.limiter))
		h.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Message": "pong",
			})
		})
	}

	u := app.Group("/user")
	{
		u.GET("/get", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Message": "user",
			})
		})
	}

	// #[handler/shop/]
	shopGroup := app.Group("/shop")
	{
		shopGroup.GET("/order", shop.OrderHandler(r.svcCtx))
	}

	return app
}

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
		// 懶得寫
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
