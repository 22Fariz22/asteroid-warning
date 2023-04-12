package asteroid

import (
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
)

type NasaWebAPI interface {
	GetNextDateWebAPI(l logger.Interface, dates []string) (int, error)
}
