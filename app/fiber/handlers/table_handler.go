package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/app/fiber/middlewares"
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
)

type TableHandler struct {
	service    service.ITableService
	middleware *middlewares.AuthMiddleware
}

func (h *TableHandler) setupRoutes(app *fiber.App) {
	tables := app.Group("/api/tables")
	tables.Use(h.middleware.CheckJWT)
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
// @Param id path string true "Table ID"
// @Tags Tables
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Table
// @Failure default {object} domain.AppError
func (h *TableHandler) findTable(c *fiber.Ctx) error {
	id := c.Params("id")
	table, err := h.service.Find(id)
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
// @Param id path string true "Table ID"
// @Param dto body dto.UpdateTableRequest true "Table UpdateTableRequest"
// @Tags Tables
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Table
// @Failure default {object} domain.AppError
func (h *TableHandler) updateTable(c *fiber.Ctx) error {
	id := c.Params("id")
	dto := new(dto.UpdateTableRequest)
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	table, err := h.service.Update(id, dto)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(table)
}

// @Router /api/tables/{id} [DELETE]
// @Security ApiKeyAuth
// @Param id path string true "Table ID"
// @Tags Tables
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Table
// @Failure default {object} domain.AppError
func (h *TableHandler) deleteTable(c *fiber.Ctx) error {
	id := c.Params("id")
	table, err := h.service.Delete(id)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(table)
}

// @Router /api/tables/{id}/tickets [POST]
// @Security ApiKeyAuth
// @Param id path string true "Table ID"
// @Tags Tables
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Table
// @Failure default {object} domain.AppError
func (h *TableHandler) saveTableTicket(c *fiber.Ctx) error {
	id := c.Params("id")
	user := c.Locals("session").(*domain.User)
	table, err := h.service.CreateTicket(id, &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(table)
}

func NewTableHandler(service service.ITableService, middleware *middlewares.AuthMiddleware, app *fiber.App) *TableHandler {
	th := TableHandler{service, middleware}
	th.setupRoutes(app)
	return &th
}
