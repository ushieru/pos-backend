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

func (h *UserHandler) listUsers(c *fiber.Ctx) error {
	users, err := h.service.List()
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(users)
}

func (h *UserHandler) findUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, err := h.service.Find(uint(id))
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(user)
}

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
