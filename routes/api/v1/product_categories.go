package api_v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/database"
	"github.com/ushieru/pos/models"
	"github.com/ushieru/pos/models/errors"
)

// @Router /api/v1/products/{productId}/categories/{categoryId} [POST]
// @Security ApiKeyAuth
// @Param productId path int true "Product ID"
// @Param categoryId path int true "Category ID"
// @Tags Product Category
// @Produce json
// @Success 200 {array} models.Category
// @Failure 0 {object} models_errors.ErrorResponse
func postCategoryProduct(c *fiber.Ctx) error {
	categoryId := c.Params("categoryId")
	productId := c.Params("productId")
	var category models.Category
	database.DBConnection.First(&category, categoryId)
	if category.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Category not found",
			fmt.Sprintf("Categoria con id: %s no encontrado", categoryId),
			"",
		))
	}
	var product models.Product
	database.DBConnection.First(&product, productId)
	if product.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Product not found",
			fmt.Sprintf("Producto con id: %s no encontrado", productId),
			"",
		))
	}
	err := database.DBConnection.Model(&category).Association("Products").Append(&product)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusOK)
}

// @Router /api/v1/products/{productId}/categories/{categoryId} [DELETE]
// @Security ApiKeyAuth
// @Param productId path int true "Product ID"
// @Param categoryId path int true "Category ID"
// @Tags Product Category
// @Produce json
// @Success 200 {array} models.Category
// @Failure 0 {object} models_errors.ErrorResponse
func deleteCategoryProduct(c *fiber.Ctx) error {
	categoryId := c.Params("categoryId")
	productId := c.Params("productId")
	var category models.Category
	database.DBConnection.First(&category, categoryId)
	if category.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Category not found",
			fmt.Sprintf("Categoria con id: %s no encontrado", categoryId),
			"",
		))
	}
	var product models.Product
	database.DBConnection.First(&product, productId)
	if product.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Product not found",
			fmt.Sprintf("Producto con id: %s no encontrado", productId),
			"",
		))
	}
	err := database.DBConnection.Model(&category).Association("Products").Delete(&product)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusOK)
}
