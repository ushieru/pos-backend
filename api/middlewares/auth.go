package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/ushieru/pos/api/database"
	"github.com/ushieru/pos/api/models"
	"github.com/ushieru/pos/api/utils"
)

func AuthMiddleware() fiber.Handler {
	return auth
}

func auth(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sessionId := uint(claims[utils.SessionParamUserId].(float64))
	var userModel models.User
	database.DBConnection.Preload("Account").First(&userModel, sessionId)
	c.Locals("session", userModel)
	return c.Next()
}
