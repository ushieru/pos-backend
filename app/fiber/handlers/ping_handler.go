package handler

import (
	"github.com/gofiber/fiber/v2"
)

type PingHandler struct{}

func (h *PingHandler) SetupRoutes(app *fiber.App) {
	app.Get("/ping", h.ping)
}

func (h *PingHandler) ping(c *fiber.Ctx) error {
	return c.SendString("\x0A")
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}
