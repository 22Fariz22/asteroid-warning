package asteroid

import (
	"github.com/22Fariz22/asteroid-warning/internal/entity"
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
)

type UseCase interface {
	GetNextDateUC(l logger.Interface, dates []string) (int, error)
	SaveAsteroidsUC(l logger.Interface, asteroids entity.NeoCounts) error
}
