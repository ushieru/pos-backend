package main_cmd

import (
	"github.com/spf13/cobra"
	main_api "github.com/ushieru/pos/api"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the web server (127.0.0.1:8080)",
	Run: func(command *cobra.Command, args []string) {
		main_api.InitServer()
	},
}
