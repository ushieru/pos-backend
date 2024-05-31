package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/ushieru/pos/app/fiber/swagger"
)

type SwaggerHandler struct{}

func (h *SwaggerHandler) setupRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}

func NewSwaggerHandler(app *fiber.App) *SwaggerHandler {
	sh := new(SwaggerHandler)
	sh.setupRoutes(app)
	return sh
}
