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

func RegisterHandlers(svc *svc.Services) *gin.Engine {
	r := gin.Default()
	app := &App{
		Router: []IRouter{
			// add here
			NewShopRouter(svc),
			NewUserRouter(svc),
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
