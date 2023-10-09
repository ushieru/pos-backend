package api_v1

import (
	"net"

	"github.com/gofiber/fiber/v2"
)

// @Router /api/v1/info [GET]
// @Tags Info
// @Produce plain
// @Success 200 {string} string "IPv4"
// @Failure 0 {object} models_errors.ErrorResponse
func getInfoRequest(c *fiber.Ctx) error {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)
	return c.SendString(localAddr.IP.String())
}