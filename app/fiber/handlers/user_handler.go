package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/app/fiber/middlewares"
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
)

type UserHandler struct {
	service    service.IUserService
	middleware *middlewares.AuthMiddleware
}

func (h *UserHandler) setupRoutes(app *fiber.App) {
	middlewareJustAdmins := middlewares.NewCheckRollMiddleware(domain.Admin)
	users := app.Group("/api/users")
	users.Use(h.middleware.CheckJWT)
	users.Get("/", h.listUsers)
	users.Get("/:id", h.findUser)
	users.Post("/", middlewareJustAdmins.CheckRol, h.saveUser)
	users.Put("/:id", middlewareJustAdmins.CheckRol, h.updateUser)
	users.Delete("/:id", middlewareJustAdmins.CheckRol, h.deleteUser)
}

// @Router /api/users [GET]
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

// @Router /api/users/{id} [GET]
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Tags User
// @Accepts json
// @Produce json
// @Success 200 {object} domain.User
// @Failure default {object} domain.AppError
func (h *UserHandler) findUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.Find(id)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(user)
}

// @Router /api/users [POST]
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

// @Router /api/users/{id} [PUT]
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Param dto body dto.UpdateUserRequest true "User UpdateUserRequest"
// @Tags User
// @Accepts json
// @Produce json
// @Success 200 {array} domain.User
// @Failure default {object} domain.AppError
func (h *UserHandler) updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	dto := new(dto.UpdateUserRequest)
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	user := c.Locals("session").(*domain.User)
	user, err := h.service.Update(id, dto, &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(user)
}

// @Router /api/users/{id} [DELETE]
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Tags User
// @Accepts json
// @Produce json
// @Success 200 {array} domain.User
// @Failure default {object} domain.AppError
func (h *UserHandler) deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := c.Locals("session").(*domain.User)
	user, err := h.service.Delete(id, &user.Account)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(user)
}

func NewUserHandler(service service.IUserService, middleware *middlewares.AuthMiddleware, app *fiber.App) *UserHandler {
	uh := UserHandler{service, middleware}
	uh.setupRoutes(app)
	return &uh
}
