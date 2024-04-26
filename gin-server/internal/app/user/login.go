package user

import (
	"context"
	"net/http"

	"free/gin-server/internal/svc"

	grpc_server "free/proto/grpc-server"

	"github.com/gin-gonic/gin"
)

type Login struct {
	svcCtx *svc.ServiceContext
}

func NewLogin(svcCtx *svc.ServiceContext) *Login {
	return &Login{
		svcCtx: svcCtx,
	}
}

func (l *Login) Login(g *gin.Context) {
	s, err := l.svcCtx.LoginRpc.Login(context.Background(), &grpc_server.LogingRequest{
		Username: "ooo",
		Passwd:   "oooo",
	})
	if err != nil {
		g.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}
	g.JSON(http.StatusOK, gin.H{
		"is login :": s.IsLogin,
	})
}

/*
type LoginStrategy interface {
	Login(c *gin.Context) (string, error)
}

type SuperCoLoginStrategy struct{}

func (sc *SuperCoLoginStrategy) Login(c *gin.Context) (string, error) {
	return "sc", nil
}

type CoLoginStrategy struct{}

func (co *CoLoginStrategy) Login(c *gin.Context) (string, error) {
	return "co", nil
}

type SuperAgentLoginStrategy struct{}

func (sa *SuperAgentLoginStrategy) Login(c *gin.Context) (string, error) {
	return "sa", nil
}

type AgentLoginStrategy struct{}

func (ag *AgentLoginStrategy) Login(c *gin.Context) (string, error) {
	return "ag", nil
}

type MemberLoginStrategy struct{}

func (mem *MemberLoginStrategy) Login(c *gin.Context) (string, error) {
	return "mem", nil
}

type LoginHandler struct {
	strategy LoginStrategy
}

func NewLoginHandler(strategy LoginStrategy) *LoginHandler {
	return &LoginHandler{
		strategy: strategy,
	}
}

func (h *LoginHandler) Handle(c *gin.Context) {
	data, err := h.strategy.Login(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

// 各階層
type LoginHandlerFactory interface {
	Create() *LoginHandler
}

// 大股東
type SuperCoLoginHandlerFactory struct{}

func (f *SuperCoLoginHandlerFactory) Create() *LoginHandler {
	return NewLoginHandler(&SuperCoLoginStrategy{})
}

// 股東
type CoLoginHandlerFactory struct{}

func (f *CoLoginHandlerFactory) Create() *LoginHandler {
	return NewLoginHandler(&CoLoginStrategy{})
}

// 總代理
type SuperAgentHandlerFactory struct{}

func (f *SuperAgentHandlerFactory) Create() *LoginHandler {
	return NewLoginHandler(&SuperAgentLoginStrategy{})
}

// 代理商
type AgentHandlerFactory struct{}

func (f *AgentHandlerFactory) Create() *LoginHandler {
	return NewLoginHandler(&AgentLoginStrategy{})
}

// Member
type MemberHandlerFactory struct{}

func (f *MemberHandlerFactory) Create() *LoginHandler {
	return NewLoginHandler(&MemberLoginStrategy{})
}
*/
