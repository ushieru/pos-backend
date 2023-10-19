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
	products := app.Group("/products")
	products.Get("/", h.listProducts)
	products.Get("/:id", h.findProduct)
	products.Post("/", h.saveProduct)
	products.Post("/:id/categories/:categoryId", h.addCategory)
	products.Put("/:id", h.updateProduct)
	products.Delete("/:id", h.deleteProduct)
}

func (h *ProductHandler) listProducts(c *fiber.Ctx) error {
	products, err := h.service.List()
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(products)
}

func (h *ProductHandler) findProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	product, err := h.service.Find(uint(id))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(product)
}

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

func (h *ProductHandler) deleteProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	product, err := h.service.Delete(uint(id))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(product)
}

func (h *ProductHandler) addCategory(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	categoryId, _ := c.ParamsInt("categoryId")
	product, err := h.service.AddCategory(uint(id), uint(categoryId))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(product)
}

func NewProductHandler(service service.IProductService) *ProductHandler {
	return &ProductHandler{service}
}
