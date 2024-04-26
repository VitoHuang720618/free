package app

import (
	"free/ws-server/internal/svc"
	"net/http"

	"github.com/gorilla/websocket"
)

type ConnectionManager struct {
	svcCtx   *svc.ServiceContext
	upgrader websocket.Upgrader
}

func NewConnectionManager(svcCtx *svc.ServiceContext) *ConnectionManager {
	return &ConnectionManager{
		svcCtx: svcCtx,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}
