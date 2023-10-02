package main

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/ushieru/pos/database"
	"github.com/ushieru/pos/middlewares"
	"github.com/ushieru/pos/routes"
	"log"

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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Point Of Sale")
	})
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/ping", routes.GetPingRequest)
	app.Get("/info", routes.GetInfoRequest)
	routes.SetupAuthRoutes(app)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("super secret word")},
	}), middlewares.AuthMiddleware())
	routes.SetupUserRoutes(app)
	routes.SetupCategoriesRoutes(app)
	routes.SetupProductsRoutes(app)
	routes.SetupTicketsRoutes(app)
	routes.SetupTableRoutes(app)
	routes.SetupPaymentsRoutes(app)
	app.Use(routes.RouteNotFound)
}
