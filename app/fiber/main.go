package fiber_app

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	docs "github.com/ushieru/pos/app/fiber/swagger"
	"github.com/ushieru/pos/domain"
)

// @Title Point Of Sale API
// @Version 1.0
// @Description Point Of Sale - Total Tools
// @Host localhost:8080
// @SecurityDefinitions.basic BasicAuth
// @In header
// @Name Authorization
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @BasePath /
func NewFiberApp(config *FiberAppConfig) *FiberApp {
	docs.SwaggerInfo.Host = fmt.Sprintf("127.0.0.1:%d", config.Port)
	app := fiber.New(fiber.Config{
		ServerHeader: "Point of Sale",
		AppName:      "Point of Sale v0.0.1",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return c.Status(code).JSON(domain.AppError{Code: code, Message: e.Error()})
		},
	})
	return &FiberApp{app, *config}
}

func Init(f *FiberApp) error {
	f.App.Use(cors.New())
	f.App.Use(func(c *fiber.Ctx) error {
		c.Locals("port", f.Config.Port)
		c.Locals("secret", f.Config.Secret)
		return c.Next()
	})
	f.App.Static("/", "public")
	return nil
}

func (f *FiberApp) Stop() error {
	return f.App.Shutdown()
}

func NewDefaultFiberAppConfig() *FiberAppConfig {
	return &FiberAppConfig{DatabaseName: "pos.db", Port: 8080, Secret: "supersecretword"}
}

type FiberApp struct {
	App    *fiber.App
	Config FiberAppConfig
}

type FiberAppConfig struct {
	DatabaseName   string
	Port           int
	Secret         string
	DatabaseLogger string
}
