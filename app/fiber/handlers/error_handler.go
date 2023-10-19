package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/domain"
)

var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	return c.Status(code).JSON(domain.AppError{Code: code, Message: e.Error()})
}
