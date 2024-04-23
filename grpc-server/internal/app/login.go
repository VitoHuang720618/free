package app

import (
	"free/grpc-server/internal/svc"
	grpc_server "free/proto/grpc-server"
)

type LoginLogic struct {
	svc *svc.Services
}

func NewLoginLogic(svc *svc.Services) *LoginLogic {
	return &LoginLogic{
		svc: svc,
	}
}

func (l *LoginLogic) Login(in *grpc_server.LogingRequest) (*grpc_server.LoginResponse, error) {
	return &grpc_server.LoginResponse{
		IsLogin: true,
	}, nil
}
