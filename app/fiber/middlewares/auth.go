package middlewares

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	fiber_app "github.com/ushieru/pos/app/fiber"
	"github.com/ushieru/pos/service"
)

type AuthMiddleware struct {
	service service.IUserService
}

func (h *AuthMiddleware) setupMiddleware(app *fiber.App) {
	app.Use(h.checkJWT, h.jwtToSession)
}

func (s *AuthMiddleware) checkJWT(c *fiber.Ctx) error {
	secret := c.Locals("secret").(string)
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
		ContextKey: "jwtToken",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, "No autorizado")
		},
	})(c)
}

func (s *AuthMiddleware) jwtToSession(c *fiber.Ctx) error {
	jwtToken := c.Locals("jwtToken").(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	userId := uint(claims["SessionParamUserId"].(float64))
	user, err := s.service.Find(userId)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	c.Locals("session", user)
	return c.Next()
}

func NewAuthMiddleware(service service.IUserService, fa *fiber_app.FiberApp) *AuthMiddleware {
	am := AuthMiddleware{service}
	am.setupMiddleware(fa.App)
	return &am
}
