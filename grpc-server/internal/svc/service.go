package svc

import "micro/grpc-server/internal/config"

type Services struct {
	Config config.Config
}

func NewServices(c config.Config) *Services {
	return &Services{
		Config: c,
	}
}
