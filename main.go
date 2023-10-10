package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/ushieru/pos/database"
	"github.com/ushieru/pos/routes"
	"github.com/ushieru/pos/routes/api"
	api_v1 "github.com/ushieru/pos/routes/api/v1"

	_ "github.com/ushieru/pos/docs"
)

// @Title Point Of Sale API
// @Version 1.0
// @Description Point Of Sale - Total Tools
// @Host localhost:8080
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @BasePath /
func main() {
	app := fiber.New()
	database.InitDatabase()
	setupRoutes(app)
	log.Fatal(app.Listen(":8080"))
}

func setupRoutes(app *fiber.App) {
	app.Use(cors.New())

	app.Static("/", "./public") 
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/ping", api.GetPingRequest)
	api_v1.SetupApiV1(app)
	
	app.Use(routes.RouteNotFound)
}
