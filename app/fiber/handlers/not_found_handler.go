package handler

import (
	"github.com/gofiber/fiber/v2"
)

type NotFoundHandler struct{}

func (h *NotFoundHandler) SetupRoutes(app *fiber.App) {
	app.Use(h.notFound)
}

func (h *NotFoundHandler) notFound(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotFound, "Ruta no encontrada")
}

func NewNotFoundHandler() *NotFoundHandler {
	return &NotFoundHandler{}
}
