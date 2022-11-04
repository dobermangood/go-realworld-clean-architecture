package fiber

import (
	"github.com/dobermangood/go-realworld-clean-architecture/config"
	"github.com/dobermangood/go-realworld-clean-architecture/internal/controller/http/fiber/handlers/article"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(
	app *fiber.App,
	conf config.Config,
	articleUseCase article.UseCase,
) {
	apiRouter := app.Group("/api")
	article.InitRoutes(apiRouter, articleUseCase)

	// middlewares.RegisterSwagger(app, conf.Server.SwaggerUrl)
	// middlewares.RegisterCheckAuthToken(app, []byte(conf.JwtSecret))

	// handlers.RegisterRegistryRoutes(app, registryUseCase)
	// handlers.RegisterMonitoringRoutes(app, monitoringUseCase)

}
