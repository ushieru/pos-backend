package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
)

type TableHandler struct {
	service service.ITableService
}

func (h *TableHandler) SetupRoutes(app *fiber.App) {
	tables := app.Group("/api/tables")
	tables.Get("/", h.listTables)
	tables.Get("/:id", h.findTable)
	tables.Post("/", h.saveTable)
	tables.Post("/:id/tickets", h.saveTableTicket)
	tables.Put("/:id", h.updateTable)
	tables.Delete("/:id", h.deleteTable)
}

// @Router /api/tables [GET]
// @Security ApiKeyAuth
// @Tags Tables
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Table
// @Failure default {object} domain.AppError
func (h *TableHandler) listTables(c *fiber.Ctx) error {
	tables, err := h.service.List()
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(tables)
}

// @Router /api/tables/{id} [GET]
// @Security ApiKeyAuth
// @Param id path int true "Table ID"
// @Tags Tables
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Table
// @Failure default {object} domain.AppError
func (h *TableHandler) findTable(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	table, err := h.service.Find(uint(id))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(table)
}

// @Router /api/tables [POST]
// @Security ApiKeyAuth
// @Param dto body dto.CreateTableRequest true "Table CreateTableRequest"
// @Tags Tables
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Table
// @Failure default {object} domain.AppError
func (h *TableHandler) saveTable(c *fiber.Ctx) error {
	dto := new(dto.CreateTableRequest)
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	table, err := h.service.Save(dto)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(table)
}

// @Router /api/tables/{id} [PUT]
// @Security ApiKeyAuth
// @Param id path int true "Table ID"
// @Param dto body dto.UpdateTableRequest true "Table UpdateTableRequest"
// @Tags Tables
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Table
// @Failure default {object} domain.AppError
func (h *TableHandler) updateTable(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	dto := new(dto.UpdateTableRequest)
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	table, err := h.service.Update(uint(id), dto)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(table)
}

// @Router /api/tables/{id} [DELETE]
// @Security ApiKeyAuth
// @Param id path int true "Table ID"
// @Tags Tables
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Table
// @Failure default {object} domain.AppError
func (h *TableHandler) deleteTable(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	table, err := h.service.Delete(uint(id))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(table)
}

// @Router /api/tables/{id}/tickets [POST]
// @Security ApiKeyAuth
// @Param id path int true "Table ID"
// @Tags Tables
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Table
// @Failure default {object} domain.AppError
func (h *TableHandler) saveTableTicket(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user := c.Locals("session").(*domain.User)
	table, err := h.service.CreateTicket(uint(id), &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(table)
}

func NewTableHandler(service service.ITableService) *TableHandler {
	return &TableHandler{service}
}
