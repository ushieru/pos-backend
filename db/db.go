package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDB(loggerLevel, connection string) *gorm.DB {
	var databaseLogger logger.Interface
	switch loggerLevel {
	case "silent":
		databaseLogger = logger.Default.LogMode(logger.Silent)
	case "info":
		databaseLogger = logger.Default.LogMode(logger.Info)
	case "error":
		databaseLogger = logger.Default.LogMode(logger.Error)
	default:
		databaseLogger = logger.Default.LogMode(logger.Silent)
	}
	database, err := gorm.Open(sqlite.Open(connection),
		&gorm.Config{Logger: databaseLogger},
	)
	if err != nil {
		panic("[Database Connection] failed to connect")
	}
	return database
}
