package app

import (
	"context"
	"free/grpc-server/internal/svc"
	grpc_server "free/proto/grpc-server"
)

// All implementations must embed UnimplementedHealthServer
// 可以參考xxx.pb.go 檔案中的上述文字的下方的interface來實作
type HealthServer struct {
	svcCtx *svc.ServiceContext
	grpc_server.UnimplementedHealthServer
}

func NewHealthServer(svcCtx *svc.ServiceContext) *HealthServer {
	return &HealthServer{
		svcCtx: svcCtx,
	}
}

func (l *HealthServer) Ping(ctx context.Context, in *grpc_server.PingRequest) (*grpc_server.PongResponse, error) {
	return &grpc_server.PongResponse{
		Message: "Pong",
	}, nil
}
