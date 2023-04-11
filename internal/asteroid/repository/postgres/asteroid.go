package postgres

import "github.com/22Fariz22/asteroid-warning/pkg/postgres"

type AsteroidRepository struct {
	*postgres.Postgres
}

func NewAsteroidRepository(db *postgres.Postgres) *AsteroidRepository {
	return &AsteroidRepository{db}
}

//here will some methods for db
