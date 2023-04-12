package asteroid

import (
	"context"
	"github.com/22Fariz22/asteroid-warning/internal/entity"
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
)

type UseCase interface {
	GetNextDateUC(ctx context.Context, l logger.Interface, dates []string) (int64, error)
	SaveAsteroidsUC(l logger.Interface, asteroids *entity.NeoCounts) error
}
