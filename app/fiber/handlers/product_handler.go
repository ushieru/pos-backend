package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
)

type ProductHandler struct {
	service service.IProductService
}

func (h *ProductHandler) SetupRoutes(app *fiber.App) {
	products := app.Group("/api/products")
	products.Get("/", h.listProducts)
	products.Get("/categories/:id", h.listProductsByCategoryId)
	products.Get("/:id", h.findProduct)
	products.Post("/", h.saveProduct)
	products.Post("/:id/categories/:categoryId", h.addCategory)
	products.Delete("/:id/categories/:categoryId", h.deleteCategory)
	products.Put("/:id", h.updateProduct)
	products.Delete("/:id", h.deleteProduct)
}

// @Router /api/products [GET]
// @Security ApiKeyAuth
// @Tags Product
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Product
// @Failure default {object} domain.AppError
func (h *ProductHandler) listProducts(c *fiber.Ctx) error {
	products, err := h.service.List()
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(products)
}

// @Router /api/products/categories/{id} [GET]
// @Security ApiKeyAuth
// @Param id path int true "Category Id"
// @Tags Product
// @Accepts json
// @Produce json
// @Success 200 {array} domain.Product
// @Failure default {object} domain.AppError
func (h *ProductHandler) listProductsByCategoryId(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	products, err := h.service.ListByCategoryId(uint(id))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(products)
}

// @Router /api/products/{id} [GET]
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Tags Product
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Product
// @Failure default {object} domain.AppError
func (h *ProductHandler) findProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	product, err := h.service.Find(uint(id))
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
	id, _ := c.ParamsInt("id")
	dto := new(dto.UpsertProductRequest)
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	product, err := h.service.Update(uint(id), dto)
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
	id, _ := c.ParamsInt("id")
	product, err := h.service.Delete(uint(id))
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
	id, _ := c.ParamsInt("id")
	categoryId, _ := c.ParamsInt("categoryId")
	product, err := h.service.AddCategory(uint(id), uint(categoryId))
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
	id, _ := c.ParamsInt("id")
	categoryId, _ := c.ParamsInt("categoryId")
	product, err := h.service.DeleteCategory(uint(id), uint(categoryId))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(product)
}

func NewProductHandler(service service.IProductService) *ProductHandler {
	return &ProductHandler{service}
}
