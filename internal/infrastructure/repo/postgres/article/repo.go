package article

import "github.com/dobermangood/go-realworld-clean-architecture/pkg/repo/postgres"

type Repo struct {
	pg *postgres.Postgres
}

func New(pg *postgres.Postgres) *Repo {
	return &Repo{
		pg: pg,
	}
}
