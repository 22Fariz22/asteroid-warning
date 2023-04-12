package http

import (
	"fmt"
	"github.com/22Fariz22/asteroid-warning/internal/asteroid"
	"github.com/22Fariz22/asteroid-warning/internal/entity"
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Handler struct {
	useCase asteroid.UseCase
	l       logger.Interface
}

func NewHandler(useCase asteroid.UseCase, l logger.Interface) *Handler {
	l.Info("create new handller")

	return &Handler{
		useCase: useCase,
		l:       l,
	}
}

// SaveAsteroids POST "/neo/count"
func (h *Handler) SaveAsteroids(c *gin.Context) {
	h.l.Info("handler SaveAsteroids")

	inp := new(entity.NeoCounts)
	if err := c.BindJSON(inp); err != nil {
		h.l.Error("failed unmarshal InputAsteroid:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err := h.useCase.SaveAsteroidsUC(h.l, *inp)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	fmt.Println("inp: ", inp)
	fmt.Println("refl inp: ", reflect.TypeOf(inp))

	c.Status(http.StatusOK)
}

// GetAsteroids GET method  "/neo/count{dates}"
// отдает ближайщую дату приближения,если есть такая среди переданных параметрах dates в адресе
func (h *Handler) GetNextDate(c *gin.Context) {
	h.l.Info("handler GetAsteroids")

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	dates := c.QueryArray("dates")
	fmt.Println("dates:", dates)

	counts, err := h.useCase.GetNextDateUC(h.l, dates)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	fmt.Println("nearest date: ", counts)

	//вернуть колтчество астероидов count
	c.Status(http.StatusOK)
}
