package postgres

import (
	"context"

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

func (a *AsteroidRepository) SaveAsteroidsRepo(l logger.Interface, asteroids *entity.NeoCounts) error {
	var existData string

	for i, _ := range asteroids.NeoCounts {
		dataNeo := asteroids.NeoCounts[i].Date
		countNEO := asteroids.NeoCounts[i].Count

		_ = a.Pool.QueryRow(context.Background(), `SELECT date FROM neo_counts where date = $1;`, dataNeo).Scan(&existData)

		if existData == dataNeo {
			_, err := a.Pool.Exec(context.Background(), `UPDATE neo_counts SET count = $1 where date=$2;`,
				countNEO, dataNeo)
			if err != nil {
				l.Error("failed update ", err)
			}
		} else {
			_, err := a.Pool.Exec(context.Background(), `INSERT INTO neo_counts (date, count) VALUES ($1,$2)`,
				dataNeo, countNEO)
			if err != nil {
				l.Error("failed insert ", err)
			}
		}

	}

	return nil
}
