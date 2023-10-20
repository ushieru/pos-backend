package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
)

type UserHandler struct {
	service service.IUserService
}

func (h *UserHandler) SetupRoutes(app *fiber.App) {
	users := app.Group("/users")
	users.Get("/", h.listUsers)
	users.Get("/:id", h.findUser)
	users.Post("/", h.saveUser)
	users.Put("/:id", h.updateUser)
	users.Delete("/:id", h.deleteUser)
}

// @Router /users [GET]
// @Security ApiKeyAuth
// @Tags User
// @Accepts json
// @Produce json
// @Success 200 {array} domain.User
// @Failure default {object} domain.AppError
func (h *UserHandler) listUsers(c *fiber.Ctx) error {
	users, err := h.service.List()
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(users)
}

// @Router /users/{id} [GET]
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Tags User
// @Accepts json
// @Produce json
// @Success 200 {object} domain.User
// @Failure default {object} domain.AppError
func (h *UserHandler) findUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, err := h.service.Find(uint(id))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(user)
}

// @Router /users [POST]
// @Security ApiKeyAuth
// @Param dto body dto.CreateUserRequest true "User CreateUserRequest"
// @Tags User
// @Accepts json
// @Produce json
// @Success 200 {array} domain.User
// @Failure default {object} domain.AppError
func (h *UserHandler) saveUser(c *fiber.Ctx) error {
	dto := new(dto.CreateUserRequest)
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	user := c.Locals("session").(*domain.User)
	user, err := h.service.Save(dto, &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(user)
}

// @Router /users/{id} [PUT]
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Param dto body dto.UpdateUserRequest true "User UpdateUserRequest"
// @Tags User
// @Accepts json
// @Produce json
// @Success 200 {array} domain.User
// @Failure default {object} domain.AppError
func (h *UserHandler) updateUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	dto := new(dto.UpdateUserRequest)
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	user := c.Locals("session").(*domain.User)
	user, err := h.service.Update(uint(id), dto, &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(user)
}

// @Router /users/{id} [DELETE]
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Tags User
// @Accepts json
// @Produce json
// @Success 200 {array} domain.User
// @Failure default {object} domain.AppError
func (h *UserHandler) deleteUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user := c.Locals("session").(*domain.User)
	user, err := h.service.Delete(uint(id), &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(user)
}

func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{service}
}
