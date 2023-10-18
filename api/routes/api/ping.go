package api

import "github.com/gofiber/fiber/v2"

// @Router /ping [GET]
// @Tags Ping
// @Produce plain
// @Success 200 {string} string "pong"
// @Failure 0 {object} models_errors.ErrorResponse
func GetPingRequest(c *fiber.Ctx) error {
	return c.SendString("pong")
}
