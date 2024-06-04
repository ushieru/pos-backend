package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/app/fiber/middlewares"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
)

type CategoryHandler struct {
	service    service.ICategoryService
	middleware *middlewares.AuthMiddleware
}

func (h *CategoryHandler) setupRoutes(app *fiber.App) {
	categories := app.Group("/api/categories")
	categories.Get("/", h.listCategories)
	categories.Get("/:id", h.findCategory)
	categories.Use(h.middleware.CheckJWT)
	categories.Post("/", h.saveCategory)
	categories.Put("/:id", h.updateCategory)
	categories.Delete("/:id", h.deleteCategory)
}

// @Router /api/categories [GET]
// @Tags Category
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Category
// @Failure default {object} domain.AppError
func (h *CategoryHandler) listCategories(c *fiber.Ctx) error {
	searchCriteriaQueryRequest := new(dto.SearchCriteriaQueryRequest)
	if err := c.QueryParser(searchCriteriaQueryRequest); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Error parse query")
	}
	withProducts := c.QueryBool("withProducts")
	categories, err := h.service.List(searchCriteriaQueryRequest, withProducts)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(categories)
}

// @Router /api/categories/{id} [GET]
// @Param id path string true "Category ID"
// @Tags Category
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Category
// @Failure default {object} domain.AppError
func (h *CategoryHandler) findCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	categories, err := h.service.Find(id)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(categories)
}

// @Router /api/categories [POST]
// @Security ApiKeyAuth
// @Param dto body dto.UpsertCategoryRequest true "Category UpsertCategoryRequest"
// @Tags Category
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Category
// @Failure default {object} domain.AppError
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

// @Router /api/categories/{id} [PUT]
// @Security ApiKeyAuth
// @Param id path string true "Category ID"
// @Param dto body dto.UpsertCategoryRequest true "Category UpsertCategoryRequest"
// @Tags Category
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Category
// @Failure default {object} domain.AppError
func (h *CategoryHandler) updateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	dto := new(dto.UpsertCategoryRequest)
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	category, err := h.service.Update(id, dto)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(category)
}

// @Router /api/categories/{id} [DELETE]
// @Security ApiKeyAuth
// @Param id path string true "Category ID"
// @Tags Category
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Category
// @Failure default {object} domain.AppError
func (h *CategoryHandler) deleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	category, err := h.service.Delete(id)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(category)
}

func NewCategoryHandler(service service.ICategoryService, middleware *middlewares.AuthMiddleware, app *fiber.App) *CategoryHandler {
	ch := CategoryHandler{service, middleware}
	ch.setupRoutes(app)
	return &ch
}
