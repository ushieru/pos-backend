package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ushieru/pos/app/fiber"
	"github.com/ushieru/pos/db"
	"github.com/ushieru/pos/repository"
	"github.com/ushieru/pos/service"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the web server",
	Run: func(command *cobra.Command, args []string) {
		config := getConfig()
		fiberApp := fiber_app.NewFiberApp(config)
		database := db.GetDB(config.DatabaseLogger, config.DatabaseName)
		userRepository := repository.NewUserGormRepository(database)
		categoryRepository := repository.NewCategoryGormRepository(database)
		productRepository := repository.NewProductGormRepository(database)
		ticketRepository := repository.NewTicketGormRepository(database)
		tableRepository := repository.NewTableGormRepository(database)

		userService := service.NewUserService(userRepository)
		categoryService := service.NewCategoryService(categoryRepository)
		productService := service.NewProductService(productRepository, categoryRepository)
		tableService := service.NewTableService(tableRepository)
		ticketService := service.NewTicketService(ticketRepository, tableRepository, productRepository)

		fiberApp.Init(&fiber_app.FiberAppServices{
			UserService:     userService,
			CategoryService: categoryService,
			ProductService:  productService,
			TableService:    tableService,
			TicketService:   ticketService,
		})
	},
}

func getConfig() *fiber_app.FiberAppConfig {
	viper.SetConfigName("pos-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	fiberConfig := fiber_app.NewDefaultFiberAppConfig()
	port := viper.GetInt("host.port")
	databaseName := viper.GetString("database.name")
	databaseLogger := viper.GetString("database.logger")
	secret := viper.GetString("jwt.secret")
	if len(fmt.Sprint(port)) == 4 {
		fiberConfig.Port = port
	}
	if databaseName != "" {
		fiberConfig.DatabaseName = databaseName
	}
	if secret != "" {
		fiberConfig.Secret = secret
	}
	fiberConfig.DatabaseLogger = databaseLogger

	return fiberConfig
}
