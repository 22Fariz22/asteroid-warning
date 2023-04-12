package http

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/22Fariz22/asteroid-warning/internal/asteroid"
	"github.com/22Fariz22/asteroid-warning/internal/entity"
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
	"github.com/gin-gonic/gin"
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

// SaveAsteroids POST "/neo/count"
func (h *Handler) SaveAsteroids(c *gin.Context) {
	inp := new(entity.NeoCounts)
	if err := c.BindJSON(inp); err != nil {
		h.l.Error("failed unmarshal InputAsteroid:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err := h.useCase.SaveAsteroidsUC(h.l, inp)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

// GetAsteroids GET method  "/neo/count{dates}"
// отдает ближайщую дату приближения,если есть такая среди переданных параметрах dates в адресе
func (h *Handler) GetNextDate(c *gin.Context) {
	//таймер на 10 секунд
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	// update gin request context
	c.Request = c.Request.WithContext(ctx)

	dates := c.QueryArray("dates")

	counts, err := h.useCase.GetNextDateUC(ctx, h.l, dates)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	//вернуть колтчество астероидов count
	c.Data(http.StatusOK, strconv.FormatInt(counts, 10), []byte(strconv.FormatInt(counts, 10)))
}
