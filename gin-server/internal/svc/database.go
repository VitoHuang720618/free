package svc

import (
	"fmt"
	"free/gin-server/internal/config"

	// "free/grpc-server/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	GameDB struct {
		Master *gorm.DB
		Slave  *gorm.DB
	}
}

func NewMysql(c config.Config) *Mysql {
	gm, err := gorm.Open(mysql.Open(c.GameDB.Master), &gorm.Config{})
	if err != nil {
		fmt.Println("Game db master connect error !! ：", err.Error())
	}

	gs, err := gorm.Open(mysql.Open(c.GameDB.Slave), &gorm.Config{})
	if err != nil {
		fmt.Println("Game db slave connect error !! ：", err.Error())
	}
	return &Mysql{
		GameDB: struct {
			Master *gorm.DB
			Slave  *gorm.DB
		}{
			Master: gm,
			Slave:  gs,
		},
	}
}
