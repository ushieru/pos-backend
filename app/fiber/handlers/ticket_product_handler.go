package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/app/fiber/middlewares"
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
)

type TicketProductHandler struct {
	service    service.ITicketProductService
	middleware *middlewares.AuthMiddleware
}

func (h *TicketProductHandler) setupRoutes(app *fiber.App) {
	ticketProducts := app.Group("/api/ticket-products")
	ticketProducts.Use(h.middleware.CheckJWT)
	ticketProducts.Get("/", h.listTicketProducts)
	ticketProducts.Put("/:id/in-preparation", h.InPreparation)
	ticketProducts.Put("/:id/prepared", h.Prepared)
	ticketProducts.Put("/:id/paid", h.Paid)
}

// @Router /api/ticket-products [GET]
// @Security ApiKeyAuth
// @Tags TicketProducts
// @Accepts json
// @Produce json
// @Success 200 {array} []domain.TicketProduct
// @Failure default {object} domain.AppError
func (h *TicketProductHandler) listTicketProducts(c *fiber.Ctx) error {
	searchCriteriaQueryRequest := new(dto.SearchCriteriaQueryRequest)
	if err := c.QueryParser(searchCriteriaQueryRequest); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Error parse query")
	}
	productionCenter, err := h.service.List(searchCriteriaQueryRequest)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(productionCenter)
}

// @Router /api/ticket-products/{id}/in-preparation [PUT]
// @Security ApiKeyAuth
// @Param id path string true "TicketProduct ID"
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

// @Router /api/ticket-products/{id}/prepared [PUT]
// @Security ApiKeyAuth
// @Param id path string true "TicketProduct ID"
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

// @Router /api/ticket-products/{id}/paid [PUT]
// @Security ApiKeyAuth
// @Param id path string true "TicketProduct ID"
// @Tags TicketProducts
// @Accepts json
// @Produce json
// @Success 200 {object} domain.TicketProduct
// @Failure default {object} domain.AppError
func (h *TicketProductHandler) Paid(c *fiber.Ctx) error {
	user := c.Locals("session").(*domain.User)
	id := c.Params("id")
	ticket, err := h.service.Paid(id, &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

func NewTicketProductHandler(service service.ITicketProductService, middleware *middlewares.AuthMiddleware, app *fiber.App) *TicketProductHandler {
	th := TicketProductHandler{service, middleware}
	th.setupRoutes(app)
	return &th
}
