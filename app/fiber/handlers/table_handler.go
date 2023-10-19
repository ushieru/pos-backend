package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
)

type TableHandler struct {
	service service.ITableService
}

func (h *TableHandler) SetupRoutes(app *fiber.App) {
	tables := app.Group("/tables")
	tables.Get("/", h.listTables)
	tables.Get("/:id", h.findTable)
	tables.Post("/", h.saveTable)
	tables.Put("/:id", h.updateTable)
	tables.Delete("/:id", h.deleteTable)
}

func (h *TableHandler) listTables(c *fiber.Ctx) error {
	tables, err := h.service.List()
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(tables)
}

func (h *TableHandler) findTable(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	table, err := h.service.Find(uint(id))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(table)
}

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

func (h *TableHandler) deleteTable(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	table, err := h.service.Delete(uint(id))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(table)
}

func NewTableHandler(service service.ITableService) *TableHandler {
	return &TableHandler{service}
}
