package server

import (
	"context"

	"free/grpc-server/internal/app"
	"free/grpc-server/internal/svc"

	grpc_server "free/proto/grpc-server"
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
