package asteroid

import (
	"github.com/22Fariz22/asteroid-warning/internal/entity"
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
)

type AsteroidRepository interface {
	GetNextDateRepo(l logger.Interface, dates []string) (*entity.Asteroid, error)
	SaveAsteroidsRepo(l logger.Interface, asteroids entity.NeoCounts) error
}
