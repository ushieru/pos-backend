package api_v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/database"
	"github.com/ushieru/pos/models"
	models_errors "github.com/ushieru/pos/models/errors"
)

func setupTicketsRoutes(app fiber.Router) {
	tickets := app.Group("/tickets")
	tickets.Get("/", getTicket)
	tickets.Get("/:id", getTicketById)
	tickets.Post("/", postTicket)
	tickets.Post("/:id/products/:productId", addProductToTicket)
	tickets.Delete("/:id/products/:productId", deleteProductToTicket)
	tickets.Delete("/:id", deleteTicket)
}

// @Router /api/v1/tickets [GET]
// @Security ApiKeyAuth
// @Param mine query boolean false "Only your tickets"
// @Param onlyOpen query boolean false "Only open tickets"
// @Tags Ticket
// @Produce json
// @Success 200 {array} models.Ticket
// @Failure 0 {object} models_errors.ErrorResponse
func getTicket(c *fiber.Ctx) error {
	mine := c.QueryBool("mine")
	onlyOpen := c.QueryBool("onlyOpen")

	var ticket []models.Ticket
	if mine {
		session := c.Locals("session").(models.User)
		if onlyOpen {
			database.DBConnection.Preload("TicketProducts").
				InnerJoins("Account").Find(&ticket, "account_id = ? AND ticket_status = ?", session.Account.ID, models.TicketOpen)
			return c.JSON(ticket)
		}
		database.DBConnection.Preload("TicketProducts").
			InnerJoins("Account").Find(&ticket, "account_id = ?", session.Account.ID)
		return c.JSON(ticket)
	}
	if onlyOpen {
		database.DBConnection.Preload("TicketProducts").
			InnerJoins("Account").Find(&ticket, "ticket_status = ?", models.TicketOpen)
		return c.JSON(ticket)
	}
	database.DBConnection.Preload("TicketProducts").InnerJoins("Account").Find(&ticket)
	return c.JSON(ticket)
}

// @Router /api/v1/tickets/{id} [GET]
// @Security ApiKeyAuth
// @Param id path int true "Ticket ID"
// @Tags Ticket
// @Produce json
// @Success 200 {object} models.Ticket
// @Failure 0 {object} models_errors.ErrorResponse
func getTicketById(c *fiber.Ctx) error {
	id := c.Params("id")
	var ticket models.Ticket
	database.DBConnection.Preload("TicketProducts").InnerJoins("Account").First(&ticket, id)
	if ticket.AccountID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"Ticket not found",
				fmt.Sprintf("Ticket con id: %s no encontrado", id),
				"",
			))
	}
	return c.JSON(ticket)
}

// @Router /api/v1/tickets [POST]
// @Security ApiKeyAuth
// @Tags Ticket
// @Produce json
// @Success 200 {object} models.Ticket
// @Failure 0 {object} models_errors.ErrorResponse
func postTicket(c *fiber.Ctx) error {
	session := c.Locals("session").(models.User)
	ticket := models.Ticket{
		TicketStatus: models.TicketOpen,
		Account:      session.Account,
	}
	database.DBConnection.Create(&ticket)
	return c.JSON(ticket)
}

// @Router /api/v1/tickets/{ticketId}/products/{productId} [POST]
// @Security ApiKeyAuth
// @Param ticketId path int true "Ticket ID"
// @Param productId path int true "Product ID"
// @Tags Ticket
// @Produce json
// @Success 200 {object} models.Ticket
// @Failure 0 {object} models_errors.ErrorResponse
func addProductToTicket(c *fiber.Ctx) error {
	session := c.Locals("session").(models.User)
	ticketId := c.Params("id")
	productId := c.Params("productId")
	var ticket models.Ticket
	database.DBConnection.First(&ticket, ticketId)
	if ticket.AccountID != session.Account.ID {
		return c.Status(fiber.StatusNotFound).JSON(
			models_errors.NewErrorResponse(
				"Ticket not found",
				fmt.Sprintf("Ticket con id: %s no encontrado", ticketId),
				"",
			))
	}
	if ticket.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(
			models_errors.NewErrorResponse(
				"Ticket not found",
				fmt.Sprintf("Ticket con id: %s no encontrado", ticketId),
				"",
			))
	}
	var product models.Product
	database.DBConnection.First(&product, productId)
	if product.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(
			models_errors.NewErrorResponse(
				"Product not found",
				fmt.Sprintf("Producto con id: %s no encontrado", productId),
				"",
			))
	}
	var ticketProduct models.TicketProduct
	database.DBConnection.First(&ticketProduct, "id = ? AND ticket_id = ?", productId, ticketId)
	if ticketProduct.ID == 0 {
		ticketProduct := models.TicketProduct{
			Product:  product,
			Quantity: 1,
		}
		database.DBConnection.Model(&ticket).Association("TicketProducts").Append(&ticketProduct)
	}
	if ticketProduct.ID != 0 {
		ticketProduct.Quantity = ticketProduct.Quantity + 1
		database.DBConnection.Save(&ticketProduct)
	}
	database.DBConnection.Preload("TicketProducts").Preload("Account").First(&ticket, ticketId)
	total := 0.0
	for _, productTicket := range ticket.TicketProducts {
		total += productTicket.Product.Price * float64(productTicket.Quantity)
	}
	ticket.Total = total
	database.DBConnection.Save(&ticket)
	return c.JSON(ticket)
}

// @Router /api/v1/tickets/{ticketId}/products/{productId} [DELETE]
// @Security ApiKeyAuth
// @Param ticketId path int true "Ticket ID"
// @Param productId path int true "Product ID"
// @Tags Ticket
// @Produce json
// @Success 200 {object} models.Ticket
// @Failure 0 {object} models_errors.ErrorResponse
func deleteProductToTicket(c *fiber.Ctx) error {
	session := c.Locals("session").(models.User)
	ticketId := c.Params("id")
	productId := c.Params("productId")
	var ticket models.Ticket
	database.DBConnection.First(&ticket, ticketId)
	if ticket.AccountID != session.Account.ID {
		return c.Status(fiber.StatusNotFound).JSON(
			models_errors.NewErrorResponse(
				"Ticket not found",
				fmt.Sprintf("Ticket con id: %s no encontrado", ticketId),
				"",
			))
	}
	if ticket.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"Ticket not found",
				fmt.Sprintf("Ticket con id: %s no encontrado", ticketId),
				"",
			))
	}
	var product models.Product
	database.DBConnection.First(&product, productId)
	if product.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"Product not found",
				fmt.Sprintf("Producto con id: %s no encontrado", productId),
				"",
			))
	}
	var ticketProduct models.TicketProduct
	database.DBConnection.First(&ticketProduct, "id = ? AND ticket_id = ?", productId, ticketId)
	if ticketProduct.ID == 0 {
		return c.JSON(ticket)
	}
	if ticketProduct.Quantity == 1 {
		database.DBConnection.Model(&ticket).Association("TicketProducts").Delete(&ticketProduct)
	}
	if ticketProduct.Quantity > 1 {
		ticketProduct.Quantity = ticketProduct.Quantity - 1
		database.DBConnection.Save(&ticketProduct)
	}
	database.DBConnection.Preload("TicketProducts").Preload("Account").First(&ticket, ticketId)
	total := 0.0
	for _, productTicket := range ticket.TicketProducts {
		total += productTicket.Product.Price * float64(productTicket.Quantity)
	}
	ticket.Total = total
	database.DBConnection.Save(&ticket)
	return c.JSON(ticket)
}

// @Router /api/v1/tickets/{id} [DELETE]
// @Security ApiKeyAuth
// @Param id path int true "Ticket ID"
// @Tags Ticket
// @Produce json
// @Success 200 {object} models.Ticket
// @Failure 0 {object} models_errors.ErrorResponse
func deleteTicket(c *fiber.Ctx) error {
	id := c.Params("id")
	var ticket models.Ticket
	database.DBConnection.Preload("TicketProducts").First(&ticket,
		"id = ? AND ticket_status = ?", id, models.TicketOpen)
	if ticket.AccountID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"Ticket not found",
				fmt.Sprintf("Ticket con id: %s no encontrado", id),
				"",
			))
	}
	if len(ticket.TicketProducts) != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"Ticket not found",
				"El ticket no puede ser eliminado porque no esta vacio",
				"",
			))
	}
	var table models.Table
	database.DBConnection.First(&table, "ticket_id = ?", id)
	if table.ID != 0 {
		database.DBConnection.Model(&table).Association("Account").Clear()
		database.DBConnection.Model(&table).Association("Ticket").Clear()
	}
	ticket.TicketStatus = models.TicketClose
	database.DBConnection.Save(&ticket)
	database.DBConnection.Delete(&ticket)
	return c.SendStatus(fiber.StatusOK)
}
