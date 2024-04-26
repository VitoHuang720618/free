package svc

import "free/ws-server/internal/config"

// import "free/gin-server/internal/config"

type Services struct {
	Config config.Config
}

func NewServices(c config.Config) *Services {
	return &Services{
		Config: c,
	}
}
