package svc

import (
	"fmt"

	"micro/gin-server/internal/config"

	"go-micro.dev/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Services struct {
	Config config.Config
	// for grpc client
	Micro  micro.Service
	GameDB struct {
		Master *gorm.DB
		Slave  *gorm.DB
	}
}

func NewServices(c config.Config) *Services {
	// example mysql init
	gm, err := gorm.Open(mysql.Open(c.GameDB.Master), &gorm.Config{})
	if err != nil {
		fmt.Println("Game db master connect error !! ：", err.Error())
	}
	gs, err := gorm.Open(mysql.Open(c.GameDB.Slave), &gorm.Config{})
	if err != nil {
		fmt.Println("Game db slave connect error !! ：", err.Error())
	}

	// redis init
	//
	//
	return &Services{
		Config: c,
		Micro:  micro.NewService(micro.Name("greeter")),
		GameDB: struct {
			Master *gorm.DB
			Slave  *gorm.DB
		}{
			Master: gm,
			Slave:  gs,
		},
	}
}
