package app

import (
	"context"
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

func (l *LoginServer) Login(ctx context.Context, in *grpc_server.LogingRequest) (*grpc_server.LoginResponse, error) {
	return &grpc_server.LoginResponse{
		IsLogin: true,
	}, nil
}
