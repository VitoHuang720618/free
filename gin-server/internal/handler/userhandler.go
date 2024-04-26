package handler

import (
	"free/gin-server/internal/app/user"
	"free/gin-server/internal/svc"

	"github.com/gin-gonic/gin"
)

type User struct {
	svcCtx *svc.ServiceContext
}

func NewUserHnadler(svcCtx *svc.ServiceContext) *User {
	return &User{
		svcCtx: svcCtx,
	}
}

func (l *User) Register(g *gin.Engine) {
	group := g.Group("/user")
	{
		login := user.NewLoginlogic(&gin.Context{}, l.svcCtx)
		group.GET("/login", login.Login)

		//v := user.NewLogout(l.svc)
		//
		//group.GET("/logut", v.Logout)
		logout := user.NewLogout(&gin.Context{}, l.svcCtx)
		group.GET("/logout", logout.Logout)
	}
}
