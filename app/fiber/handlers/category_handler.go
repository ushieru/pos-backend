package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
)

type CategoryHandler struct {
	service service.ICategoryService
}

func (h *CategoryHandler) SetupRoutes(app *fiber.App) {
	categories := app.Group("/categories")
	categories.Get("/", h.listCategories)
	categories.Get("/:id", h.findCategory)
	categories.Post("/", h.saveCategory)
	categories.Put("/:id", h.updateCategory)
	categories.Delete("/:id", h.deleteCategory)
}

// @Router /categories [GET]
// @Security ApiKeyAuth
// @Tags Category
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Category
// @Failure 0 {object} domain.AppError
func (h *CategoryHandler) listCategories(c *fiber.Ctx) error {
	categories, err := h.service.List()
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(categories)
}

// @Router /categories/{id} [GET]
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Tags Category
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Category
// @Failure 0 {object} domain.AppError
func (h *CategoryHandler) findCategory(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	categories, err := h.service.Find(uint(id))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(categories)
}

// @Router /categories [POST]
// @Security ApiKeyAuth
// @Param dto body dto.UpsertCategoryRequest true "Category UpsertCategoryRequest"
// @Tags Category
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Category
// @Failure 0 {object} domain.AppError
func (h *CategoryHandler) saveCategory(c *fiber.Ctx) error {
	dto := new(dto.UpsertCategoryRequest)
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	category, err := h.service.Save(dto)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(category)
}

// @Router /categories/{id} [PUT]
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Param dto body dto.UpsertCategoryRequest true "Category UpsertCategoryRequest"
// @Tags Category
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Category
// @Failure 0 {object} domain.AppError
func (h *CategoryHandler) updateCategory(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	dto := new(dto.UpsertCategoryRequest)
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	category, err := h.service.Update(uint(id), dto)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(category)
}

// @Router /categories/{id} [DELETE]
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Tags Category
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Category
// @Failure 0 {object} domain.AppError
func (h *CategoryHandler) deleteCategory(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	category, err := h.service.Delete(uint(id))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(category)
}

func NewCategoryHandler(service service.ICategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}
