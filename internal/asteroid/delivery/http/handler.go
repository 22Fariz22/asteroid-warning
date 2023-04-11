package http

import (
	"github.com/22Fariz22/asteroid-warning/internal/asteroid"
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	useCase asteroid.UseCase
	l       logger.Interface
}

func NewHandler(useCase asteroid.UseCase, l logger.Interface) *Handler {

	return &Handler{
		useCase: useCase,
		l:       l,
	}
}

// SaveAsteroids
func (h *Handler) SaveAsteroids(c *gin.Context) {
	h.l.Info("here handler SaveAsteroids")

	c.Status(http.StatusOK)

}

// GetAsteroids
func (h *Handler) GetAsteroids(c *gin.Context) {
	h.l.Info("here handler GetAsteroids")

	c.Status(http.StatusOK)
}
