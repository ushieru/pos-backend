package fiber_app

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ushieru/pos/app/fiber/handlers"
	"github.com/ushieru/pos/app/fiber/middlewares"
	docs "github.com/ushieru/pos/app/fiber/swagger"
	"github.com/ushieru/pos/repository"
	"github.com/ushieru/pos/service"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		ErrorHandler: handler.DefaultErrorHandler,
	})
	return &FiberApp{app, *config}
}

func (f *FiberApp) Run() error {
	database, err := gorm.Open(sqlite.Open(f.config.DatabaseName),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)
	if err != nil {
		panic("[Database Connection] failed to connect")
	}

	notFoundHandler := handler.NewNotFoundHandler()
	pingHandler := handler.NewPingHandler()
	swaggerHandler := handler.NewSwaggerHandler()
	userRepository := repository.NewUserGormRepository(database)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(userService)
	authMiddleware := middlewares.NewAuthMiddleware(userService)
	categoryRepository := repository.NewCategoryGormRepository(database)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productRepository := repository.NewProductGormRepository(database)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	tableRepository := repository.NewTableGormRepository(database)
	tableService := service.NewTableService(tableRepository)
	tableHandler := handler.NewTableHandler(tableService)

	f.app.Use(cors.New())
	f.app.Use(func(c *fiber.Ctx) error {
		c.Locals("port", f.config.Port)
		c.Locals("secret", f.config.Secret)
		return c.Next()
	})
	f.app.Static("/", "public")
	swaggerHandler.SetupRoutes(f.app)
	pingHandler.SetupRoutes(f.app)
	authHandler.SetupRoutes(f.app)
	authMiddleware.SetupMiddleware(f.app)
	userHandler.SetupRoutes(f.app)
	categoryHandler.SetupRoutes(f.app)
	productHandler.SetupRoutes(f.app)
	tableHandler.SetupRoutes(f.app)
	notFoundHandler.SetupRoutes(f.app)
	return f.app.Listen(fmt.Sprintf(":%d", f.config.Port))
}

func NewDefaultFiberAppConfig() *FiberAppConfig {
	return &FiberAppConfig{"pos.db", 8080, "supersecretword"}
}

type FiberApp struct {
	app    *fiber.App
	config FiberAppConfig
}

type FiberAppConfig struct {
	DatabaseName string
	Port         int
	Secret       string
}
