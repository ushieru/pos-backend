package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/app/fiber"
)

type PingHandler struct{}

func (h *PingHandler) setupRoutes(app *fiber.App) {
	app.Get("/ping", h.ping)
}

// @Router /ping [GET]
// @Tags Ping
// @Success 200
func (h *PingHandler) ping(c *fiber.Ctx) error {
	return c.SendString("\x0A")
}

func NewPingHandler(fa *fiber_app.FiberApp) *PingHandler {
	ph := new(PingHandler)
	ph.setupRoutes(fa.App)
	return ph
}
