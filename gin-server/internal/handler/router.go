package handler

import (
	"free/gin-server/internal/svc"

	"github.com/gin-gonic/gin"
)

type IRouter interface {
	Register(r *gin.Engine)
}

type App struct {
	Router []IRouter
}

func RegisterHandlers(svcCtx *svc.ServiceContext) *gin.Engine {
	r := gin.Default()
	app := &App{
		Router: []IRouter{
			// add here
			NewShopRouter(svcCtx),
			NewUserRouter(svcCtx),
		},
	}
	app.registerHandlers(r)
	return r
}

func (app *App) registerHandlers(r *gin.Engine) {
	for _, handler := range app.Router {
		handler.Register(r)
	}
}
