package postgres

import (
	"github.com/22Fariz22/asteroid-warning/internal/entity"
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
	"github.com/22Fariz22/asteroid-warning/pkg/postgres"
)

type AsteroidRepository struct {
	*postgres.Postgres
}

func NewAsteroidRepository(db *postgres.Postgres) *AsteroidRepository {
	return &AsteroidRepository{db}
}

func (a *AsteroidRepository) SaveAsteroidsRepo(l logger.Interface, asteroids entity.NeoCounts) error {
	l.Info("SaveAsteroidsRepo.")
	return nil
}
