package middlewares

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ushieru/pos/service"
)

type AuthMiddleware struct {
	service service.IUserService
}

func (s *AuthMiddleware) CheckJWT(c *fiber.Ctx) error {
	secret := c.Locals("secret").(string)
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
		ContextKey: "jwtToken",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, "No autorizado")
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			jwtToken := c.Locals("jwtToken").(*jwt.Token)
			claims := jwtToken.Claims.(jwt.MapClaims)
			userId := claims["SessionParamUserId"].(string)
			user, err := s.service.Find(userId)
			if err != nil {
				return fiber.NewError(err.Code, err.Message)
			}
			c.Locals("session", user)
			return c.Next()
		},
	})(c)
}

func NewAuthMiddleware(service service.IUserService) *AuthMiddleware {
	return &AuthMiddleware{service}
}
