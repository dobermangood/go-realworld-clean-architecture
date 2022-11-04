package main

import (
	"context"
	"log"

	"github.com/dobermangood/go-realworld-clean-architecture/config"
	_app "github.com/dobermangood/go-realworld-clean-architecture/internal/app"

	_fiberServer "github.com/dobermangood/go-realworld-clean-architecture/pkg/httpserver/fiber"
	_postgres "github.com/dobermangood/go-realworld-clean-architecture/pkg/repo/postgres"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatal("config loading err: ", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	postgres, err := _postgres.New(ctx, conf.Postgres)
	if err != nil {
		log.Fatal("postgres connection error: ", err)
	}

	err = _postgres.RunMigrations(conf.Postgres)
	if err != nil {
		log.Fatal("postgres run migrations error: ", err)
	}

	server := _fiberServer.New()
	app := _app.New(ctx, conf, server, postgres)

	err = app.Run()
	if err != nil {
		log.Fatal("server err: ", err)
	}
}
