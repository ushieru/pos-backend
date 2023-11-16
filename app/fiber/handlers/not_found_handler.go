package handler

import (
	"github.com/gofiber/fiber/v2"
	fiber_app "github.com/ushieru/pos/app/fiber"
)

type NotFoundHandler struct{}

func (h *NotFoundHandler) setupRoutes(app *fiber.App) {
	app.Use(h.notFound)
}

func (h *NotFoundHandler) notFound(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotFound, "Ruta no encontrada")
}

func NewNotFoundHandler(fa *fiber_app.FiberApp) *NotFoundHandler {
	nfh := new(NotFoundHandler)
	nfh.setupRoutes(fa.App)
	return nfh
}
