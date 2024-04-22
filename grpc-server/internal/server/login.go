package server

import (
	"context"

	"micro/grpc-server/internal/app"
	"micro/grpc-server/internal/svc"

	grpc_server "micro/proto/grpc-server"
)

type LoginServer struct {
	svc *svc.Services
	grpc_server.UnimplementedLoginServer
}

func NewLogin(svc *svc.Services) *LoginServer {
	return &LoginServer{
		svc: svc,
	}
}

func (h *LoginServer) Login(ctx context.Context, in *grpc_server.LogingRequest) (*grpc_server.LoginResponse, error) {
	l := app.NewLoginLogic(h.svc)
	return l.Login(in)
}
