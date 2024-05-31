package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/app/fiber/middlewares"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
)

type ProductHandler struct {
	service    service.IProductService
	middleware *middlewares.AuthMiddleware
}

func (h *ProductHandler) setupRoutes(app *fiber.App) {
	products := app.Group("/api/products")
	products.Get("/", h.listProducts)
	products.Get("/categories/:id", h.listProductsByCategoryId)
	products.Get("/:id", h.findProduct)
	products.Use(h.middleware.CheckJWT)
	products.Post("/", h.saveProduct)
	products.Post("/:id/categories/:categoryId", h.addCategory)
	products.Delete("/:id/categories/:categoryId", h.deleteCategory)
	products.Put("/:id", h.updateProduct)
	products.Delete("/:id", h.deleteProduct)
}

// @Router /api/products [GET]
// @Param criteria query dto.SearchCriteriaQueryRequest false "Criteria filter"
// @Tags Product
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Product
// @Failure default {object} domain.AppError
func (h *ProductHandler) listProducts(c *fiber.Ctx) error {
	searchCriteriaQueryRequest := new(dto.SearchCriteriaQueryRequest)
	if err := c.QueryParser(searchCriteriaQueryRequest); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Error parse query")
	}
	withCategories := c.QueryBool("withCategories")
	products, err := h.service.List(searchCriteriaQueryRequest, withCategories)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(products)
}

// @Router /api/products/categories/{id} [GET]
// @Param id path int true "Category Id"
// @Tags Product
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Product
// @Failure default {object} domain.AppError
func (h *ProductHandler) listProductsByCategoryId(c *fiber.Ctx) error {
	id := c.Params("id")
	searchCriteriaQueryRequest := new(dto.SearchCriteriaQueryRequest)
	if err := c.QueryParser(searchCriteriaQueryRequest); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Error parse query")
	}
	products, err := h.service.ListByCategoryId(id, searchCriteriaQueryRequest)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(products)
}

// @Router /api/products/{id} [GET]
// @Param id path int true "Product ID"
// @Tags Product
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Product
// @Failure default {object} domain.AppError
func (h *ProductHandler) findProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := h.service.Find(id)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(product)
}

// @Router /api/products [POST]
// @Security ApiKeyAuth
// @Param dto body dto.UpsertProductRequest true "Product UpsertProductRequest"
// @Tags Product
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Product
// @Failure default {object} domain.AppError
func (h *ProductHandler) saveProduct(c *fiber.Ctx) error {
	dto := new(dto.UpsertProductRequest)
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	product, err := h.service.Save(dto)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(product)
}

// @Router /api/products/{id} [PUT]
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param dto body dto.UpsertProductRequest true "Product UpsertProductRequest"
// @Tags Product
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Product
// @Failure default {object} domain.AppError
func (h *ProductHandler) updateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	dto := new(dto.UpsertProductRequest)
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	product, err := h.service.Update(id, dto)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(product)
}

// @Router /api/products/{id} [DELETE]
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Tags Product
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Product
// @Failure default {object} domain.AppError
func (h *ProductHandler) deleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := h.service.Delete(id)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(product)
}

// @Router /api/products/{id}/categories/{categoryId} [POST]
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param categoryId path int true "Category ID"
// @Tags Product
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Product
// @Failure default {object} domain.AppError
func (h *ProductHandler) addCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	categoryId := c.Params("categoryId")
	product, err := h.service.AddCategory(id, categoryId)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(product)
}

// @Router /api/products/{id}/categories/{categoryId} [DELETE]
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param categoryId path int true "Category ID"
// @Tags Product
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Product
// @Failure default {object} domain.AppError
func (h *ProductHandler) deleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	categoryId := c.Params("categoryId")
	product, err := h.service.DeleteCategory(id, categoryId)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(product)
}

func NewProductHandler(service service.IProductService, middleware *middlewares.AuthMiddleware, app *fiber.App) *ProductHandler {
	ph := ProductHandler{service, middleware}
	ph.setupRoutes(app)
	return &ph
}
