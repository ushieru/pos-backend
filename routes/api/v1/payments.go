package api_v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/database"
	"github.com/ushieru/pos/models"
	models_errors "github.com/ushieru/pos/models/errors"
)

func setupPaymentsRoutes(app fiber.Router) {
	tickets := app.Group("/payments")
	tickets.Post("/tickets/:ticketId", payTicket)
}

// @Router /api/v1/payments/tickets/{ticketId} [POST]
// @Security ApiKeyAuth
// @Param ticketId path int true "Ticket ID"
// @Tags Payments
// @Produce json
// @Success 200 {object} models.Ticket
// @Failure 0 {object} models_errors.ErrorResponse
func payTicket(c *fiber.Ctx) error {
	session := c.Locals("session").(models.User)
	ticketId := c.Params("ticketId")
	var ticket models.Ticket
	database.DBConnection.Preload("TicketProducts").Preload("Account").First(&ticket, ticketId)
	if session.Account.AccountType != models.Cashier {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(models_errors.NewErrorResponse(
			"Auth Error",
			"No estas autorizado para realizar esta accion",
			"",
		))
	}
	if ticket.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models_errors.NewErrorResponse(
			"Ticket not found",
			fmt.Sprintf("Ticket con id: %s, no encontrado", ticketId),
			"",
		))
	}
	if ticket.TicketStatus == models.TicketClose {
		return c.Status(fiber.StatusNotFound).JSON(models_errors.NewErrorResponse(
			"Ticket close",
			fmt.Sprintf("Ticket con id: %s, ya esta cerrado", ticketId),
			"",
		))
	}
	var table models.Table
	database.DBConnection.First(&table, "ticket_id = ?", ticketId)
	if table.ID != 0 {
		database.DBConnection.Model(&table).Association("Account").Clear()
		database.DBConnection.Model(&table).Association("Ticket").Clear()
	}
	ticket.TicketStatus = models.TicketClose
	database.DBConnection.Save(&ticket)
	return c.JSON(ticket)
}
