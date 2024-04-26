package app

import (
	"free/ws-server/internal/svc"
	"net/http"

	"github.com/gorilla/websocket"
)

type ConnectionManager struct {
	svc      *svc.Services
	upgrader websocket.Upgrader
}

func NewConnectionManager(svc *svc.Services) *ConnectionManager {
	return &ConnectionManager{
		svc: svc,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}
