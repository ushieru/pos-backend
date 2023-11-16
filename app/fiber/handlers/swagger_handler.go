package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	fiber_app "github.com/ushieru/pos/app/fiber"
	_ "github.com/ushieru/pos/app/fiber/swagger"
)

type SwaggerHandler struct{}

func (h *SwaggerHandler) setupRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}

func NewSwaggerHandler(fa *fiber_app.FiberApp) *SwaggerHandler {
	sh := new(SwaggerHandler)
	sh.setupRoutes(fa.App)
	return sh
}
