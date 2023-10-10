package database

import (
	"fmt"
	"gorm.io/gorm/logger"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	DBConnection *gorm.DB
)

func InitDatabase() {
	var err error
	DBConnection, err = gorm.Open(sqlite.Open("pos.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println(err)
		panic("[Database Connection] failed to connect")
	}
	fmt.Println("[Database Connection] done")
	Migrations()
	Seed()
}
