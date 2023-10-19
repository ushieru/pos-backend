package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/service"
)

type TicketHandler struct {
	service service.ITicketService
}

func (h *TicketHandler) SetupRoutes(app *fiber.App) {
	tickets := app.Group("/tickets")
	tickets.Get("/", h.listTickets)
	tickets.Get("/:id", h.findTicket)
	tickets.Post("/", h.saveTicket)
	tickets.Delete("/:id", h.deleteTicket)
}

func (h *TicketHandler) listTickets(c *fiber.Ctx) error {
	tickets, err := h.service.List()
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(tickets)
}

func (h *TicketHandler) findTicket(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	ticket, err := h.service.Find(uint(id))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

func (h *TicketHandler) saveTicket(c *fiber.Ctx) error {
	user := c.Locals("session").(*domain.User)
	ticket, err := h.service.Save(&user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

func (h *TicketHandler) deleteTicket(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	ticket, err := h.service.Delete(uint(id))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(ticket)
}

func NewTicketHandler(service service.ITicketService) *TicketHandler {
	return &TicketHandler{service}
}
