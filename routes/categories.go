package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/database"
	"github.com/ushieru/pos/models"
	"github.com/ushieru/pos/models/dto"
	"github.com/ushieru/pos/models/errors"
	"github.com/ushieru/pos/utils"
)

func SetupCategoriesRoutes(app *fiber.App) {
	categories := app.Group("/categories")
	categories.Get("/", getCategory)
	categories.Get("/:id", getCategoryById)
	categories.Post("/", postCategory)
	categories.Post("/:categoryId/products/:productId", postCategoryProduct)
	categories.Put("/:id", putCategory)
	categories.Delete("/:id", deleteCategory)
}

// @Router /categories [GET]
// @Security ApiKeyAuth
// @Tags Category
// @Produce json
// @Success 200 {array} models.Category
func getCategory(c *fiber.Ctx) error {
	var category []models.Category
	database.DBConnection.Find(&category)
	return c.JSON(category)
}

// @Router /categories/{id} [GET]
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Tags Category
// @Produce json
// @Success 200 {object} models.Category
func getCategoryById(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	database.DBConnection.First(&category, id)
	if category.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse("Category not found",
				fmt.Sprintf("Categoria con id: %s, no encontrada", id), ""))
	}
	return c.JSON(category)
}

// @Router /categories [POST]
// @Security ApiKeyAuth
// @Param category body dto.CreateCategoryDTO true "Category DTO"
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} models.Category
// @Failure 0 {object} models_errors.ErrorResponse
func postCategory(c *fiber.Ctx) error {
	categoryDTO := new(dto.CreateCategoryDTO)
	if err := c.BodyParser(categoryDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"Params error", "Params error", ""))
	}
	validationError := utils.ValidateStruct(categoryDTO)
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"Validation error",
				"Error al intentar crear nueva categoria",
				validationError.ToString(),
			))
	}
	category := new(models.Category)
	category.Name = categoryDTO.Name
	database.DBConnection.Create(&category)
	return c.JSON(category)
}

// @Router /categories/{id} [PUT]
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Param category body dto.UpdateCategoryDTO true "Category DTO"
// @Tags Category
// @Produce json
// @Success 200
// @Failure 0 {object} models_errors.ErrorResponse
func putCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	updateCategoryDTO := new(dto.UpdateCategoryDTO)
	if err := c.BodyParser(updateCategoryDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"Params error", "Params error", ""))
	}
	validationError := utils.ValidateStruct(updateCategoryDTO)
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Validation error",
			"Error al actualizar categoria",
			validationError.ToString(),
		))

	}
	category := new(models.Category)
	database.DBConnection.First(&category, id)
	if category.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Category not found",
			fmt.Sprintf("Categoria con id: %s, no encontrada", id),
			"",
		))
	}
	category.Name = updateCategoryDTO.Name
	database.DBConnection.Save(&category)
	return c.JSON(category)
}

// @Router /categories/{id} [DELETE]
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Tags Category
// @Produce json
// @Success 200
// @Failure 0 {object} models_errors.ErrorResponse
func deleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	database.DBConnection.First(&category, id)
	if category.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Category not found",
			fmt.Sprintf("Categoria con id: %s, no encontrada", id),
			"",
		))
	}
	database.DBConnection.Delete(&category)
	return c.SendStatus(fiber.StatusOK)
}
