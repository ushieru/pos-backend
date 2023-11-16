package handler

import (
	"encoding/base64"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/app/fiber"
	"github.com/ushieru/pos/service"
)

type AuthHandler struct {
	service service.IUserService
}

func (h *AuthHandler) setupRoutes(app *fiber.App) {
	app.Post("/api/auth", h.auth)
}

// @Router /api/auth [POST]
// @Tags Auth
// @Security BasicAuth
// @Accepts json
// @Produce json
// @Success 200 {object} dto.AuthUserResponse
// @Failure default {object} domain.AppError
func (h *AuthHandler) auth(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	authHeader := headers["Authorization"]
	authHeaderSlice := strings.Split(authHeader, " ")
	if len(authHeaderSlice) < 2 {
		return fiber.NewError(fiber.StatusBadRequest, "Bad request")
	}
	b64Credentials := authHeaderSlice[1]
	credentials, err := base64.StdEncoding.DecodeString(b64Credentials)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad request")
	}
	credentialsSlice := strings.Split(string(credentials), ":")
	if len(credentialsSlice) < 2 {
		return fiber.NewError(fiber.StatusBadRequest, "Bad request")
	}
	username := credentialsSlice[0]
	password := credentialsSlice[1]
	secret := c.Locals("secret").(string)
	authResponse, authErr := h.service.AuthWithCredentials(username, password, secret)
	if authErr != nil {
		return fiber.NewError(authErr.Code, authErr.Message)
	}
	return c.JSON(authResponse)
}

func NewAuthHandler(service service.IUserService, fa *fiber_app.FiberApp) *AuthHandler {
	ah := AuthHandler{service}
	ah.setupRoutes(fa.App)
	return &ah
}
