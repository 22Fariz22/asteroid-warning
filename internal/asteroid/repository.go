package asteroid

import (
	"github.com/22Fariz22/asteroid-warning/internal/entity"
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
)

type AsteroidRepository interface {
	SaveAsteroidsRepo(l logger.Interface, asteroids *entity.NeoCounts) error
}
