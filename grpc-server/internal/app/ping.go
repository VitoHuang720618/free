package app

import (
	"context"
	"free/grpc-server/internal/svc"
	grpc_server "free/proto/grpc-server"
)

type PingServer struct {
	svc *svc.Services
}

func NewPing(svc *svc.Services) *PingLogic {
	return &PingLogic{
		svc: svc,
	}
}

func (l *PingLogic) Ping(ctx context.Context, in *grpc_server.PingRequest) (*grpc_server.PongResponse, error) {
	return &grpc_server.PongResponse{
		Message: "Pong",
	}, nil
}
