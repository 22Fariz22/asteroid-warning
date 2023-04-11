package http

import (
	"github.com/22Fariz22/asteroid-warning/internal/asteroid"
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpointsOrder(router *gin.RouterGroup, uc asteroid.UseCase, l logger.Interface) {
	h := NewHandler(uc, l)

	asteroid := router.Group("/")
	{
		asteroid.POST("", h.SaveAsteroids)
		asteroid.GET("", h.GetAsteroids)
	}
}
