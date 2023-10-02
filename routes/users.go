package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/database"
	"github.com/ushieru/pos/models"
	models_errors "github.com/ushieru/pos/models/errors"
	"github.com/ushieru/pos/utils"
)

func SetupUserRoutes(app *fiber.App) {
	users := app.Group("/users")
	users.Get("/", getUser)
	users.Get("/:id", getUserById)
	users.Post("/", postUser)
	users.Put("/:id", putUser)
	users.Delete("/:id", deleteUser)
}

// @Router /users [GET]
// @Security ApiKeyAuth
// @Tags User
// @Produce json
// @Success 200 {array} models.User
// @Failure 0 {object} models_errors.ErrorResponse
func getUser(c *fiber.Ctx) error {
	var users []models.User
	database.DBConnection.Preload("Account").Find(&users)
	return c.JSON(users)
}

// @Router /users/{id} [GET]
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Tags User
// @Produce json
// @Success 200 {object} models.User
// @Failure 0 {object} models_errors.ErrorResponse
func getUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DBConnection.Preload("Account").First(&user, id)
	if user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"User not found",
				fmt.Sprintf("Usuario con id: %s no encontrado", id),
				"",
			))
	}
	return c.JSON(user)
}

func postUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"Params error", "Params error", ""))
	}
	validationError := utils.ValidateStruct(user)
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Validation error",
			"Error al actualizar categoria",
			validationError.ToString(),
		))
	}
	database.DBConnection.Create(&user)
	return c.JSON(user)
}

func putUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var userParams models.User
	if err := c.BodyParser(userParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"Params error", "Params error", ""))
	}
	validationError := utils.ValidateStruct(userParams)
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models_errors.NewErrorResponse(
			"Validation error",
			"Error al actualizar categoria",
			validationError.ToString(),
		))
	}
	user := new(models.User)
	database.DBConnection.First(&user, id)
	if user.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"User not found",
				fmt.Sprintf("Usuario con id: %s no encontrado", id),
				"",
			))
	}
	user.Name = userParams.Name
	user.Email = userParams.Email
	database.DBConnection.Save(&user)
	return c.JSON(user)
}

// @Router /users/{id} [DELETE]
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Tags User
// @Success 200
// @Failure 0 {object} models_errors.ErrorResponse
func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DBConnection.First(&user, id)
	if user.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse(
				"User not found",
				fmt.Sprintf("Usuario con id: %s no encontrado", id),
				"",
			))
	}
	database.DBConnection.Delete(&user)
	return c.SendStatus(fiber.StatusOK)
}
