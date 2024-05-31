package handler

import "github.com/gofiber/fiber/v2"

type NotFoundHandler struct{}

func (h *NotFoundHandler) setupRoutes(app *fiber.App) {
	app.Use(h.notFound)
}

func (h *NotFoundHandler) notFound(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotFound, "Ruta no encontrada")
}

func NewNotFoundHandler(app *fiber.App) *NotFoundHandler {
	nfh := new(NotFoundHandler)
	nfh.setupRoutes(app)
	return nfh
}
