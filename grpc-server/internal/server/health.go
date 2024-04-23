package server

import (
	"context"

	"free/grpc-server/internal/app"
	"free/grpc-server/internal/svc"

	grpc_server "free/proto/grpc-server"
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
