package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/domain"
)

type CheckRollMiddleware struct {
	rol domain.AccountType
}

func (s *CheckRollMiddleware) CheckRol(c *fiber.Ctx) error {
	user := c.Locals("session").(*domain.User)
	if user.Account.AccountType != s.rol {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}
	return c.Next()
}

func NewCheckRollMiddleware(rol domain.AccountType) *CheckRollMiddleware {
	crm := CheckRollMiddleware{rol}
	return &crm
}
