package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/app/fiber/middlewares"
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
)

type TicketHandler struct {
	service    service.ITicketService
	middleware *middlewares.AuthMiddleware
}

func (h *TicketHandler) setupRoutes(app *fiber.App) {
	tickets := app.Group("/api/tickets")
	tickets.Use(h.middleware.CheckJWT)
	tickets.Get("/", h.listTickets)
	tickets.Get("/:id", h.findTicket)
	tickets.Post("/", h.saveTicket)
	tickets.Delete("/:id", h.deleteTicket)
	tickets.Post("/:id/products/:productId", h.addProduct)
	tickets.Delete("/:id/products/:productId", h.deleteProduct)
	tickets.Put("/:id/pay", h.payTicket)
	tickets.Put("/:id/order", h.orderTicket)
}

// @Router /api/tickets [GET]
// @Security ApiKeyAuth
// @Param criteria query dto.SearchCriteriaQueryRequest false "Criteria filter"
// @Tags Tickets
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Ticket
// @Failure default {object} domain.AppError
func (h *TicketHandler) listTickets(c *fiber.Ctx) error {
	searchCriteriaQueryRequest := new(dto.SearchCriteriaQueryRequest)
	if err := c.QueryParser(searchCriteriaQueryRequest); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Error parse query")
	}
	tickets, err := h.service.List(searchCriteriaQueryRequest)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(tickets)
}

// @Router /api/tickets/{id} [GET]
// @Security ApiKeyAuth
// @Param id path string true "Ticket ID"
// @Tags Tickets
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Ticket
// @Failure default {object} domain.AppError
func (h *TicketHandler) findTicket(c *fiber.Ctx) error {
	id := c.Params("id")
	ticket, err := h.service.Find(id)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

// @Router /api/tickets [POST]
// @Security ApiKeyAuth
// @Tags Tickets
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Ticket
// @Failure default {object} domain.AppError
func (h *TicketHandler) saveTicket(c *fiber.Ctx) error {
	user := c.Locals("session").(*domain.User)
	ticket, err := h.service.Save(&user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

// @Router /api/tickets/{id} [DELETE]
// @Security ApiKeyAuth
// @Param id path string true "Ticket ID"
// @Tags Tickets
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Ticket
// @Failure default {object} domain.AppError
func (h *TicketHandler) deleteTicket(c *fiber.Ctx) error {
	id := c.Params("id")
	ticket, err := h.service.Delete(id)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

// @Router /api/tickets/{id}/products/{productId} [POST]
// @Security ApiKeyAuth
// @Param id path string true "Ticket ID"
// @Param productId path int true "Product ID"
// @Tags Tickets
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Ticket
// @Failure default {object} domain.AppError
func (h *TicketHandler) addProduct(c *fiber.Ctx) error {
	user := c.Locals("session").(*domain.User)
	id := c.Params("id")
	productId := c.Params("productId")
	ticket, err := h.service.AddProduct(id, productId, &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

// @Router /api/tickets/{id}/products/{productId} [DELETE]
// @Security ApiKeyAuth
// @Param id path string true "Ticket ID"
// @Param productId path int true "Product ID"
// @Tags Tickets
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Ticket
// @Failure default {object} domain.AppError
func (h *TicketHandler) deleteProduct(c *fiber.Ctx) error {
	user := c.Locals("session").(*domain.User)
	id := c.Params("id")
	productId := c.Params("productId")
	ticket, err := h.service.DeleteProduct(id, productId, &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

// @Router /api/tickets/{id}/pay [PUT]
// @Security ApiKeyAuth
// @Param id path string true "Ticket ID"
// @Tags Tickets
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Ticket
// @Failure default {object} domain.AppError
func (h *TicketHandler) payTicket(c *fiber.Ctx) error {
	user := c.Locals("session").(*domain.User)
	id := c.Params("id")
	ticket, err := h.service.PayTicket(id, &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

// @Router /api/tickets/{id}/order [PUT]
// @Security ApiKeyAuth
// @Param id path string true "Ticket ID"
// @Tags Tickets
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Ticket
// @Failure default {object} domain.AppError
func (h *TicketHandler) orderTicket(c *fiber.Ctx) error {
	user := c.Locals("session").(*domain.User)
	id := c.Params("id")
	ticket, err := h.service.OrderTicketProducts(id, &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

func NewTicketHandler(service service.ITicketService, middleware *middlewares.AuthMiddleware, app *fiber.App) *TicketHandler {
	th := TicketHandler{service, middleware}
	th.setupRoutes(app)
	return &th
}
