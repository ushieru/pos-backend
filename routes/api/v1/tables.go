package api_v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/database"
	"github.com/ushieru/pos/models"
	"github.com/ushieru/pos/models/dto"
	models_errors "github.com/ushieru/pos/models/errors"
	"github.com/ushieru/pos/utils"
)

func setupTableRoutes(app fiber.Router) {
	products := app.Group("/tables")
	products.Get("/", getTables)
	products.Get("/:id", getTableById)
	products.Post("/", createTable)
	products.Post("/:id/ticket", createTableTicket)
	products.Put("/:id", updateTable)
	products.Delete("/:id", deleteTable)
}

// @Router /api/v1/tables [GET]
// @Security ApiKeyAuth
// @Tags Table
// @Produce json
// @Success 200 {array} models.Table
// @Failure 0 {object} models_errors.ErrorResponse
func getTables(c *fiber.Ctx) error {
	var tables []models.Table
	database.DBConnection.
		Preload("Account").
		Preload("Ticket").
		Preload("Ticket.Account").
		Preload("Ticket.TicketProducts").
		Find(&tables)
	return c.JSON(tables)
}

// @Router /api/v1/tables/{id} [GET]
// @Security ApiKeyAuth
// @Param id path int true "Table ID"
// @Tags Table
// @Produce json
// @Success 200 {array} models.Table
// @Failure 0 {object} models_errors.ErrorResponse
func getTableById(c *fiber.Ctx) error {
	id := c.Params("id")
	var table models.Table
	database.DBConnection.Preload("Account").Preload("Ticket").First(&table, id)
	if table.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Table not found",
			fmt.Sprintf("Mesa con id: %s no encontrado", id),
			"",
		))
	}
	return c.JSON(table)
}

// @Router /api/v1/tables [POST]
// @Security ApiKeyAuth
// @Param createTableDto body dto.CreateTableDTO true "Create Table DTO"
// @Tags Table
// @Produce json
// @Success 200 {array} models.Table
// @Failure 0 {object} models_errors.ErrorResponse
func createTable(c *fiber.Ctx) error {
	createTableDTO := new(dto.CreateTableDTO)
	if err := c.BodyParser(createTableDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"Params error", "Params error", ""))
	}
	validationError := utils.ValidateStruct(createTableDTO)
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Validation error",
			"Error al crear mesa",
			validationError.ToString(),
		))
	}
	var tableFind models.Table
	database.DBConnection.First(&tableFind, "pos_x = ? AND pos_y = ?", createTableDTO.PosX, createTableDTO.PosY)
	if tableFind.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Table not found",
			fmt.Sprintf("Mesa con coordenadas (x:%d, y:%d) ya ocupada", createTableDTO.PosX, createTableDTO.PosY),
			"",
		))
	}
	table := new(models.Table)
	table.Name = createTableDTO.Name
	table.PosX = uint(createTableDTO.PosX)
	table.PosY = uint(createTableDTO.PosY)
	database.DBConnection.Save(&table)
	return c.JSON(table)
}

// @Router /api/v1/tables/{id} [PUT]
// @Security ApiKeyAuth
// @Param id path int true "Table id"
// @Param createTableDto body dto.CreateTableDTO true "Update Table DTO"
// @Tags Table
// @Produce json
// @Success 200 {array} models.Table
// @Failure 0 {object} models_errors.ErrorResponse
func updateTable(c *fiber.Ctx) error {
	id := c.Params("id")
	createTableDTO := new(dto.CreateTableDTO)
	if err := c.BodyParser(createTableDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"Params error", "Params error", ""))
	}
	validationError := utils.ValidateStruct(createTableDTO)
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Validation error",
			"Error al crear mesa",
			validationError.ToString(),
		))
	}
	var tableFind models.Table
	database.DBConnection.First(&tableFind, "pos_x = ? AND pos_y = ?", createTableDTO.PosX, createTableDTO.PosY)
	if tableFind.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Table not found",
			fmt.Sprintf("Mesa con coordenadas (x:%d, y:%d) ya ocupada", createTableDTO.PosX, createTableDTO.PosY),
			"",
		))
	}
	table := new(models.Table)
	database.DBConnection.First(&table, id)
	if table.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Table not found",
			fmt.Sprintf("Mesa con id: %s no encontrada", id),
			"",
		))
	}
	table.Name = createTableDTO.Name
	table.PosX = uint(createTableDTO.PosX)
	table.PosY = uint(createTableDTO.PosY)
	database.DBConnection.Save(&table)
	return c.JSON(table)
}

// @Router /api/v1/tables/{id}/ticket [POST]
// @Security ApiKeyAuth
// @Param id path int true "Table ID"
// @Tags Table
// @Produce json
// @Success 200 {array} models.Table
// @Failure 0 {object} models_errors.ErrorResponse
func createTableTicket(c *fiber.Ctx) error {
	id := c.Params("id")
	table := new(models.Table)
	session := c.Locals("session").(models.User)
	database.DBConnection.First(&table, id)
	if table.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Table not found",
			fmt.Sprintf("Mesa con id: %s no encontrada", id),
			"",
		))
	}
	ticket := models.Ticket{
		TicketStatus: models.TicketOpen,
		Account:      session.Account,
	}
	database.DBConnection.Create(&ticket)
	table.AccountID = session.Account.UserID
	table.TicketID = ticket.ID
	database.DBConnection.Save(&table)
	database.DBConnection.Preload("Account").Preload("Ticket").Preload("Ticket.Account").First(&table, id)
	return c.JSON(table)
}

// @Router /api/v1/tables/{id} [DELETE]
// @Security ApiKeyAuth
// @Param id path int true "Table ID"
// @Tags Table
// @Produce json
// @Success 200 {array} models.Table
// @Failure 0 {object} models_errors.ErrorResponse
func deleteTable(c *fiber.Ctx) error {
	id := c.Params("id")
	table := new(models.Table)
	database.DBConnection.First(&table, id)
	if table.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Table not found",
			fmt.Sprintf("Mesa con id: %s no encontrada", id),
			"",
		))
	}
	if table.TicketID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Table is not empty",
			"La mesa aun tiene un ticket abierto",
			"",
		))
	}
	database.DBConnection.Unscoped().Delete(&table)
	return c.SendStatus(fiber.StatusOK)
}
