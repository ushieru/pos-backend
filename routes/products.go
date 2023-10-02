package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/database"
	"github.com/ushieru/pos/models"
	dto "github.com/ushieru/pos/models/dto"
	"github.com/ushieru/pos/models/errors"
	"github.com/ushieru/pos/utils"
)

func SetupProductsRoutes(app *fiber.App) {
	products := app.Group("/products")
	products.Get("/", getProduct)
	products.Get("/:id", getProductById)
	products.Get("/categories/:idCategory", getProductByCategory)
	products.Post("/:productId/categories/:categoryId", postCategoryProduct)
	products.Delete("/:productId/categories/:categoryId", deleteCategoryProduct)
	products.Post("/", postProduct)
	products.Put("/:id", putProduct)
	products.Delete("/:id", deleteProduct)
}

// @Router /products [GET]
// @Security ApiKeyAuth
// @Tags Product
// @Produce json
// @Success 200 {array} models.Product
// @Failure 0 {object} models_errors.ErrorResponse
func getProduct(c *fiber.Ctx) error {
	var product []models.Product
	database.DBConnection.Find(&product)
	return c.JSON(product)
}

// @Router /products/{id} [GET]
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Tags Product
// @Produce json
// @Success 200 {object} models.Product
// @Failure 0 {object} models_errors.ErrorResponse
func getProductById(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	database.DBConnection.Preload("Categories").First(&product, id)
	if product.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Product not found",
			fmt.Sprintf("Producto con id: %s no encontrado", id),
			"",
		))
	}
	return c.JSON(product)
}

// @Router /products/categories/{idCategory} [GET]
// @Security ApiKeyAuth
// @Param idCategory path int true "Category ID"
// @Tags Product
// @Produce json
// @Success 200 {array} models.Product
// @Failure 0 {object} models_errors.ErrorResponse
func getProductByCategory(c *fiber.Ctx) error {
	categoryId := c.Params("idCategory")
	var category models.Category
	database.DBConnection.Find(&category, categoryId)
	if category.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Category not found",
			fmt.Sprintf("Categoria con id: %s no encontrado", categoryId),
			"",
		))
	}
	database.DBConnection.Preload("Products").Find(&category)
	return c.JSON(category.Products)
}

// @Router /products [POST]
// @Security ApiKeyAuth
// @Param ProductDTO body dto.CreateProductDTO true "Product DTO"
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} models.Product
// @Failure 0 {object} models_errors.ErrorResponse
func postProduct(c *fiber.Ctx) error {
	createProductDTO := new(dto.CreateProductDTO)
	fmt.Println(string(c.Body()))
	if err := c.BodyParser(createProductDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"Params error", "Params error", ""))
	}
	validationError := utils.ValidateStruct(createProductDTO)
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Validation error",
			"Error al crear producto",
			validationError.ToString(),
		))

	}
	product := new(models.Product)
	product.Name = createProductDTO.Name
	product.Description = createProductDTO.Description
	product.Price = createProductDTO.Price
	database.DBConnection.Create(&product)
	return c.JSON(product)
}

// @Router /products/{id} [PUT]
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param ProductDTO body dto.UpdateProductDTO true "Product DTO"
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} models.Product
// @Failure 0 {object} models_errors.ErrorResponse
func putProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	updateProductDTO := new(dto.UpdateProductDTO)
	if err := c.BodyParser(updateProductDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"Params error", "Params error", ""))
	}
	validationError := utils.ValidateStruct(updateProductDTO)
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Validation error",
			"Error al actualizar producto",
			validationError.ToString(),
		))

	}
	product := new(models.Product)
	database.DBConnection.First(&product, id)
	if product.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Product not found",
			fmt.Sprintf("Producto con id: %s no encontrado", id),
			"",
		))
	}
	if updateProductDTO.Name != "" {
		product.Name = updateProductDTO.Name
	}
	if updateProductDTO.Description != "" {
		product.Description = updateProductDTO.Description
	}
	if updateProductDTO.Price != 0 {
		product.Price = updateProductDTO.Price
	}
	database.DBConnection.Save(&product)
	return c.JSON(product)
}

// @Router /products/{id} [DELETE]
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Tags Product
// @Success 200
// @Failure 0 {object} models_errors.ErrorResponse
func deleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	database.DBConnection.First(&product, id)
	if product.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Product not found",
			fmt.Sprintf("Producto con id: %s no encontrado", id),
			"",
		))
	}
	database.DBConnection.Delete(&product)
	return c.SendStatus(fiber.StatusOK)
}
