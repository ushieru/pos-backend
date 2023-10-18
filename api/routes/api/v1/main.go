package api_v1

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/api/middlewares"
)

func SetupApiV1(app *fiber.App) {
	apiV1 := app.Group("/api/v1")
	apiV1.Get("/info", getInfoRequest)
	setupAuthRoutes(apiV1)
	apiV1.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("super secret word")},
	}), middlewares.AuthMiddleware())
	setupUserRoutes(apiV1)
	setupCategoriesRoutes(apiV1)
	setupProductsRoutes(apiV1)
	setupTicketsRoutes(apiV1)
	setupTableRoutes(apiV1)
	setupPaymentsRoutes(apiV1)
}
