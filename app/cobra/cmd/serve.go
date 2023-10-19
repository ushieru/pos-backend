package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ushieru/pos/app/fiber"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the web server",
	Run: func(command *cobra.Command, args []string) {
		fiberConfig := fiber_app.NewDefaultFiberAppConfig()
		fiberApp := fiber_app.NewFiberApp(fiberConfig)
		fiberApp.Run()
	},
}
