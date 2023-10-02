package database

import (
	"fmt"

	"github.com/ushieru/pos/models"
)

func Migrations() {
	if err := DBConnection.AutoMigrate(&models.Account{}); err != nil {
		panic("[Database Migrations] Error in Account migration")
	}
	if err := DBConnection.AutoMigrate(&models.User{}); err != nil {
		panic("[Database Migrations] Error in User migration")
	}
	if err := DBConnection.AutoMigrate(&models.Product{}); err != nil {
		panic("[Database Migrations] Error in Product migration")
	}
	if err := DBConnection.AutoMigrate(&models.Category{}); err != nil {
		panic("[Database Migrations] Error in Category migration")
	}
	if err := DBConnection.AutoMigrate(&models.TicketProduct{}); err != nil {
		panic("[Database Migrations] Error in TicketProduct migration")
	}
	if err := DBConnection.AutoMigrate(&models.Ticket{}); err != nil {
		panic("[Database Migrations] Error in Ticket migration")
	}
	if err := DBConnection.AutoMigrate(&models.Table{}); err != nil {
		panic("[Database Migrations] Error in Ticket migration")
	}

	fmt.Println("[Database Migrations] done")
}
