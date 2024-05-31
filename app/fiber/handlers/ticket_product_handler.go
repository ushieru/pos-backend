package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/app/fiber/middlewares"
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/service"
)

type TicketProductHandler struct {
	service    service.ITicketProductService
	middleware *middlewares.AuthMiddleware
}

func (h *TicketProductHandler) setupRoutes(app *fiber.App) {
	ticketProducts := app.Group("/api/ticket-products")
	ticketProducts.Use(h.middleware.CheckJWT)
	ticketProducts.Put("/:id/in-preparation")
	ticketProducts.Put("/:id/prepared")
}

// @Router /api/ticket-products/{id}/in-preparation [PUT]
// @Security ApiKeyAuth
// @Param id path int true "TicketProduct ID"
// @Tags TicketProducts
// @Accepts json
// @Produce json
// @Success 200 {object} domain.TicketProduct
// @Failure default {object} domain.AppError
func (h *TicketProductHandler) InPreparation(c *fiber.Ctx) error {
	user := c.Locals("session").(*domain.User)
	id := c.Params("id")
	ticket, err := h.service.InPreparation(id, &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

// @Router /api/ticket-products/{id}/in-prepared [PUT]
// @Security ApiKeyAuth
// @Param id path int true "TicketProduct ID"
// @Tags TicketProducts
// @Accepts json
// @Produce json
// @Success 200 {object} domain.TicketProduct
// @Failure default {object} domain.AppError
func (h *TicketProductHandler) Prepared(c *fiber.Ctx) error {
	user := c.Locals("session").(*domain.User)
	id := c.Params("id")
	ticket, err := h.service.Prepared(id, &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

func NewTicketProductHandlerTicketHandler(service service.ITicketProductService, middleware *middlewares.AuthMiddleware, app *fiber.App) *TicketProductHandler {
	th := TicketProductHandler{service, middleware}
	th.setupRoutes(app)
	return &th
}
