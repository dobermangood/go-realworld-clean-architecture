package app

import (
	"context"

	"github.com/dobermangood/go-realworld-clean-architecture/config"
	"github.com/dobermangood/go-realworld-clean-architecture/pkg/repo/postgres"

	_fiberRouter "github.com/dobermangood/go-realworld-clean-architecture/internal/controller/http/fiber"
	_articleRepo "github.com/dobermangood/go-realworld-clean-architecture/internal/infrastructure/repo/postgres/article"
	_articleUseCase "github.com/dobermangood/go-realworld-clean-architecture/internal/usecase/article"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	conf       config.Config
	ctx        context.Context
	httpserver *fiber.App
	postgres   *postgres.Postgres
}

func New(
	ctx context.Context,
	conf config.Config,
	httpserver *fiber.App,
	postgres *postgres.Postgres,
) *App {
	articleRepo := _articleRepo.New(postgres)
	articleUseCase := _articleUseCase.New(articleRepo)

	_fiberRouter.InitRouter(httpserver, conf, articleUseCase)

	return &App{
		ctx:        ctx,
		conf:       conf,
		httpserver: httpserver,
		postgres:   postgres,
	}
}

func (a *App) Run() error {
	return a.httpserver.Listen(":" + a.conf.Server.Port)
}
