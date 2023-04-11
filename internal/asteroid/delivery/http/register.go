package http

import (
	"github.com/22Fariz22/asteroid-warning/internal/asteroid"
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpointsOrder(router *gin.RouterGroup, uc asteroid.UseCase, l logger.Interface) {
	h := NewHandler(uc, l)
	h.l.Info("handler RegisterHTTPEndpointsOrder")

	asteroid := router.Group("/neo/count")
	{
		asteroid.POST("/", h.SaveAsteroids)
		asteroid.GET("/", h.GetNextDate)
	}
}
