package handler

import (
	"github.com/gofiber/fiber/v2"
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

func NewPingHandler(app *fiber.App) *PingHandler {
	ph := new(PingHandler)
	ph.setupRoutes(app)
	return ph
}
