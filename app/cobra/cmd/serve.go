package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ushieru/pos/app/fiber"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the web server",
	Run: func(command *cobra.Command, args []string) {
		fiberConfig := getConfig()
		fiberApp := fiber_app.NewFiberApp(fiberConfig)
		fiberApp.Run()
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
