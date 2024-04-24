package svc

import (
	"fmt"

	"free/gin-server/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlClient struct {
	GameMaster *gorm.DB
	GameSlave  *gorm.DB
}

func NewMySqlClient(c config.Config) *MySqlClient {
	return &MySqlClient{
		GameMaster: newMySql(c.DB.GameDB.Master),
		GameSlave:  newMySql(c.DB.GameDB.Slave),
	}
}

func newMySql(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to initialize MySQL databases : ", err)
	}
	return db
}
