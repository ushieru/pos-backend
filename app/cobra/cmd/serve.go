package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ushieru/pos/app/fiber"
	handler "github.com/ushieru/pos/app/fiber/handlers"
	"github.com/ushieru/pos/app/fiber/middlewares"
	"github.com/ushieru/pos/db"
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/repository"
	"github.com/ushieru/pos/service"
	"go.uber.org/fx"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the web server",
	Run: func(command *cobra.Command, args []string) {
		fx.New(
			fx.Provide(
				getConfig,
				newFiberApp,
				db.GetDB,
				fx.Annotate(repository.NewUserGormRepository, fx.As(new(domain.IUserRepository))),
				fx.Annotate(
					repository.NewCategoryGormRepository,
					fx.As(new(domain.ICategoryRepository)),
				),
				fx.Annotate(
					repository.NewProductGormRepository,
					fx.As(new(domain.IProductRepository)),
				),
				fx.Annotate(
					repository.NewTicketGormRepository,
					fx.As(new(domain.ITicketRepository)),
				),
				fx.Annotate(repository.NewTableGormRepository, fx.As(new(domain.ITableRepository))),
				fx.Annotate(service.NewUserService, fx.As(new(service.IUserService))),
				fx.Annotate(service.NewCategoryService, fx.As(new(service.ICategoryService))),
				fx.Annotate(service.NewProductService, fx.As(new(service.IProductService))),
				fx.Annotate(service.NewTicketService, fx.As(new(service.ITicketService))),
				fx.Annotate(service.NewTableService, fx.As(new(service.ITableService))),
			),
			fx.Invoke(
				fiber_app.Init,
				handler.NewPingHandler,
				handler.NewInfoHandler,
				handler.NewSwaggerHandler,
				handler.NewAuthHandler,
				middlewares.NewAuthMiddleware,
				handler.NewUserHandler,
				handler.NewCategoryHandler,
				handler.NewProductHandler,
				handler.NewTicketHandler,
				handler.NewTableHandler,
				handler.NewNotFoundHandler,
			),
			fx.NopLogger,
		).Run()
	},
}

func newFiberApp(lc fx.Lifecycle, c *fiber_app.FiberAppConfig) *fiber_app.FiberApp {
	fiberApp := fiber_app.NewFiberApp(c)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go fiberApp.App.Listen(fmt.Sprintf(":%d", c.Port))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return fiberApp.Stop()
		},
	})
	return fiberApp
}

func getConfig() *fiber_app.FiberAppConfig {
	viper.SetConfigName("pos-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	fiberConfig := fiber_app.NewDefaultFiberAppConfig()
	port := viper.GetInt("host.port")
	databaseName := viper.GetString("database.name")
	databaseLogger := viper.GetString("database.logger")
	secret := viper.GetString("jwt.secret")
	if len(fmt.Sprint(port)) == 4 {
		fiberConfig.Port = port
	}
	if databaseName != "" {
		fiberConfig.DatabaseName = databaseName
	}
	if secret != "" {
		fiberConfig.Secret = secret
	}
	fiberConfig.DatabaseLogger = databaseLogger

	return fiberConfig
}
