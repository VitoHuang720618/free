package server

import (
	"context"

	"micro/grpc-server/internal/app"
	"micro/grpc-server/internal/svc"

	grpc_server "micro/proto/grpc-server"
)

type HealthServer struct {
	svc *svc.Services
	grpc_server.UnimplementedHealthServer
}

func NewHealth(svc *svc.Services) *HealthServer {
	return &HealthServer{
		svc: svc,
	}
}

func (h *HealthServer) Ping(ctx context.Context, in *grpc_server.PingRequest) (*grpc_server.PongResponse, error) {
	l := app.NewPingLogic(h.svc)
	return l.Ping(in)
}
