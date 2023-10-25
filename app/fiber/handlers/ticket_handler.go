package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
)

type TicketHandler struct {
	service service.ITicketService
}

func (h *TicketHandler) SetupRoutes(app *fiber.App) {
	tickets := app.Group("/api/tickets")
	tickets.Get("/", h.listTickets)
	tickets.Get("/:id", h.findTicket)
	tickets.Post("/", h.saveTicket)
	tickets.Delete("/:id", h.deleteTicket)
	tickets.Post("/:id/products/:productId", h.addProduct)
	tickets.Delete("/:id/products/:productId", h.deleteProduct)
	tickets.Put("/:id/pay", h.payTicket)
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
// @Param id path int true "Ticket ID"
// @Tags Tickets
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Ticket
// @Failure default {object} domain.AppError
func (h *TicketHandler) findTicket(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	ticket, err := h.service.Find(uint(id))
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
// @Param id path int true "Ticket ID"
// @Tags Tickets
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Ticket
// @Failure default {object} domain.AppError
func (h *TicketHandler) deleteTicket(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	ticket, err := h.service.Delete(uint(id))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

// @Router /api/tickets/{id}/products/{productId} [POST]
// @Security ApiKeyAuth
// @Param id path int true "Ticket ID"
// @Param productId path int true "Product ID"
// @Tags Tickets
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Ticket
// @Failure default {object} domain.AppError
func (h *TicketHandler) addProduct(c *fiber.Ctx) error {
	user := c.Locals("session").(*domain.User)
	id, _ := c.ParamsInt("id")
	productId, _ := c.ParamsInt("productId")
	ticket, err := h.service.AddProduct(uint(id), uint(productId), &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

// @Router /api/tickets/{id}/products/{productId} [DELETE]
// @Security ApiKeyAuth
// @Param id path int true "Ticket ID"
// @Param productId path int true "Product ID"
// @Tags Tickets
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Ticket
// @Failure default {object} domain.AppError
func (h *TicketHandler) deleteProduct(c *fiber.Ctx) error {
	user := c.Locals("session").(*domain.User)
	id, _ := c.ParamsInt("id")
	productId, _ := c.ParamsInt("productId")
	ticket, err := h.service.DeleteProduct(uint(id), uint(productId), &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

// @Router /api/tickets/{id}/pay [PUT]
// @Security ApiKeyAuth
// @Param id path int true "Ticket ID"
// @Tags Tickets
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Ticket
// @Failure default {object} domain.AppError
func (h *TicketHandler) payTicket(c *fiber.Ctx) error {
	user := c.Locals("session").(*domain.User)
	id, _ := c.ParamsInt("id")
	ticket, err := h.service.PayTicket(uint(id), &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

func NewTicketHandler(service service.ITicketService) *TicketHandler {
	return &TicketHandler{service}
}
