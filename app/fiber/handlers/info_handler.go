package handler

import (
	"net"

	"github.com/gofiber/fiber/v2"
	fiber_app "github.com/ushieru/pos/app/fiber"
	"github.com/ushieru/pos/domain"
)

type InfoHandler struct{}

func (h *InfoHandler) setupRoutes(app *fiber.App) {
	app.Get("/info", h.getInfo)
}

// @Router /info [GET]
// @Tags Info
// @Success 200 {object} domain.Info
func (h *InfoHandler) getInfo(c *fiber.Ctx) error {
	port := c.Locals("port").(int)
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := localAddr.IP.String()
	conn.Close()
	info := domain.Info{
		Ip:   ip,
		Port: uint16(port),
	}
	return c.JSON(info)
}

func NewInfoHandler(fa *fiber_app.FiberApp) *InfoHandler {
	ih := new(InfoHandler)
	ih.setupRoutes(fa.App)
	return ih
}
