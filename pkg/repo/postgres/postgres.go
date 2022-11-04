package postgres

import (
	"context"

	"github.com/dobermangood/go-realworld-clean-architecture/config"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	Pool    *pgxpool.Pool
	Builder squirrel.StatementBuilderType
}

func New(ctx context.Context, conf config.Postgres) (*Postgres, error) {
	postgresConf, err := pgxpool.ParseConfig(conf.ConnString)
	if err != nil {
		return nil, err
	}

	postgresConf.MaxConns = conf.MaxConnections
	pool, err := pgxpool.ConnectConfig(ctx, postgresConf)
	if err != nil {
		return nil, err
	}

	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	pg := &Postgres{
		Pool:    pool,
		Builder: builder,
	}

	return pg, err
}
