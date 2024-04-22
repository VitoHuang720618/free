package handler

import (
	"micro/gin-server/internal/app/user"
	"micro/gin-server/internal/svc"

	"github.com/gin-gonic/gin"
)

type User struct {
	svc *svc.Services
}

func NewUserRouter(svc *svc.Services) *User {
	return &User{
		svc: svc,
	}
}

func (l *User) Register(g *gin.Engine) {
	group := g.Group("/user")
	{
		login := user.NewLogin(l.svc)
		group.GET("/login", login.Login)

		//v := user.NewLogout(l.svc)
		//
		//group.GET("/logut", v.Logout)
		logout := user.NewLogout(l.svc)
		group.GET("/logout", logout.Logout)
	}
}
