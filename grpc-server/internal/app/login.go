package app

import (
	"context"
	"free/grpc-server/internal/svc"
	grpc_server "free/proto/grpc-server"
)

//implement 實作這些
/*
type LoginServer interface {
	Login(context.Context, *LogingRequest) (*LoginResponse, error)
	Logout(context.Context, *Empty) (*LoginResponse, error)
	CheckSession(context.Context, *LogingRequest) (*LoginResponse, error)
}
*/

type LoginServer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	grpc_server.UnimplementedLoginServer
}

func NewLoginServer(svcCtx *svc.ServiceContext) *LoginServer {
	return &LoginServer{
		svcCtx: svcCtx,
	}
}

func (l *LoginServer) Login(ctx context.Context, in *grpc_server.LogingRequest) (*grpc_server.LoginResponse, error) {
	return &grpc_server.LoginResponse{
		IsLogin: true,
	}, nil
}

func (l *LoginServer) Logout(ctx context.Context, in *grpc_server.Empty) (*grpc_server.LoginResponse, error) {
	return &grpc_server.LoginResponse{
		IsLogin: true,
	}, nil
}

func (l *LoginServer) CheckSession(ctx context.Context, in *grpc_server.LogingRequest) (*grpc_server.LoginResponse, error) {
	return &grpc_server.LoginResponse{
		IsLogin: true,
	}, nil
}
