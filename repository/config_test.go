package repository

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getInMemoryDB() (*gorm.DB, func() error) {
	database, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	db, _ := database.DB()
	return database, db.Close
}
