package database

import (
	"fmt"
	"github.com/ushieru/pos/api/models"
	"github.com/ushieru/pos/api/utils"
)

func Seed() {
	createUsersIfNotExists()
	createCategoriesIfNotExists()
	createProductsIfNotExists()
}

func createUsersIfNotExists() {
	var user models.User
	DBConnection.First(&user)
	if user.ID != 0 {
		return
	}
	password, _ := utils.HashPassword("admin")
	isActive := true
	DBConnection.Create(&models.User{
		Name:  "Admin",
		Email: "admin@email.com",
		Account: models.Account{
			Username:    "admin",
			Password:    password,
			IsActive:    &isActive,
			AccountType: models.Admin,
		},
	})
	password, _ = utils.HashPassword("cashier")
	isActive = true
	DBConnection.Create(&models.User{
		Name:  "Cashier",
		Email: "cashier@email.com",
		Account: models.Account{
			Username:    "cashier",
			Password:    password,
			IsActive:    &isActive,
			AccountType: models.Cashier,
		},
	})
	password, _ = utils.HashPassword("waiter")
	isActive = true
	DBConnection.Create(&models.User{
		Name:  "Waiter",
		Email: "waiter@email.com",
		Account: models.Account{
			Username:    "waiter",
			Password:    password,
			IsActive:    &isActive,
			AccountType: models.Waiter,
		},
	})
	fmt.Println("[Database Seed] create admin")
}

func createCategoriesIfNotExists() {
	var category models.Category
	DBConnection.First(&category)
	if category.ID != 0 {
		return
	}
	categories := []string{"desayunos", "comidas", "bebidas"}
	for _, category := range categories {
		DBConnection.Create(&models.Category{Name: category})
	}
	fmt.Println("[Database Seed] create categories")
}

func createProductsIfNotExists() {
	var product models.Product
	DBConnection.First(&product)
	if product.ID != 0 {
		return
	}
	products := []models.Product{
		{Name: "huevos revueltos", Description: "", Price: 10},
		{Name: "chilaquiles", Description: "", Price: 15.5},
		{Name: "hotcakes", Description: "", Price: 12},
	}
	var category models.Category
	DBConnection.First(&category, 1)
	for _, product := range products {
		DBConnection.Create(&product).Association("Categories").Append(&category)
	}
	products = []models.Product{
		{Name: "hamburguesa", Description: "", Price: 10},
		{Name: "camarones", Description: "", Price: 15.5},
		{Name: "tostada de ceviiche", Description: "", Price: 12},
	}
	var category2 models.Category
	DBConnection.First(&category2, 2)
	for _, product := range products {
		DBConnection.Create(&product).Association("Categories").Append(&category2)
	}
	products = []models.Product{
		{Name: "cerveza", Description: "", Price: 10},
		{Name: "agua de naranja", Description: "", Price: 15.5},
		{Name: "horchata", Description: "", Price: 12},
	}
	var category3 models.Category
	DBConnection.First(&category3, 3)
	for _, product := range products {
		DBConnection.Create(&product).Association("Categories").Append(&category3)
	}
	fmt.Println("[Database Seed] create products")
}
