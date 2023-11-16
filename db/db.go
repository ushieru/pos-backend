package db

import (
	"github.com/glebarez/sqlite"
	fiber_app "github.com/ushieru/pos/app/fiber"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDB(f *fiber_app.FiberApp) *gorm.DB {
	var databaseLogger logger.Interface
	switch f.Config.DatabaseLogger {
	case "silent":
		databaseLogger = logger.Default.LogMode(logger.Silent)
	case "info":
		databaseLogger = logger.Default.LogMode(logger.Info)
	case "error":
		databaseLogger = logger.Default.LogMode(logger.Error)
	default:
		databaseLogger = logger.Default.LogMode(logger.Silent)
	}
	database, err := gorm.Open(sqlite.Open(f.Config.DatabaseName),
		&gorm.Config{Logger: databaseLogger},
	)
	if err != nil {
		panic("[Database Connection] failed to connect")
	}
	return database
}
