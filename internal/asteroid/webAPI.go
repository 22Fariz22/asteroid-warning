package asteroid

import (
	"context"
	"time"

	"github.com/22Fariz22/asteroid-warning/pkg/logger"
)

type NasaWebAPI interface {
	GetNextDateWebAPI(ctx context.Context, l logger.Interface, dates []time.Time) (int64, error)
}
