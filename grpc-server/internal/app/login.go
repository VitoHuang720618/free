package app

import (
	"free/grpc-server/internal/svc"
	grpc_server "free/proto/grpc-server"
)

type LoginServer struct {
	svc *svc.Services
	grpc_server.UnimplementedLoginServer
}

func NewLogin(svc *svc.Services) *LoginLogic {
	return &LoginLogic{
		svc: svc,
	}
}

func (l *LoginLogic) Login(ctx contex.Context, in *grpc_server.LogingRequest) (*grpc_server.LoginResponse, error) {
	return &grpc_server.LoginResponse{
		IsLogin: true,
	}, nil
}
