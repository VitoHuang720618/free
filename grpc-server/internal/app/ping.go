package app

import (
	"free/grpc-server/internal/svc"
	grpc_server "free/proto/grpc-server"
)

type PingLogic struct {
	svc *svc.Services
}

func NewPingLogic(svc *svc.Services) *PingLogic {
	return &PingLogic{
		svc: svc,
	}
}

func (l *PingLogic) Ping(in *grpc_server.PingRequest) (*grpc_server.PongResponse, error) {
	return &grpc_server.PongResponse{
		Message: "Pong",
	}, nil
}