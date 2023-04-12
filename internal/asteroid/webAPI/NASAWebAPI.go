package webAPI

import (
	"context"
	"time"

	"github.com/22Fariz22/asteroid-warning/pkg/logger"
	"github.com/peteretelej/nasa"
)

type WebAPI struct {
}

func NewWebAPI() *WebAPI {
	return &WebAPI{}
}

var neoL []*nasa.NeoList

func (w *WebAPI) GetNextDateWebAPI(ctx context.Context, l logger.Interface, dates []time.Time) (int64, error) {
	for i, _ := range dates {
		//добавить таймер чтобы не больше 10 сек
		//взять первый попавшийся count не ноль
		neo, err := nasa.NeoFeed(dates[i], dates[i])
		if err != nil {
			l.Error("failed nasa.NeoFeed(): ", err)
			continue
		}
		if neo.ElementCount != 0 {
			return neo.ElementCount, nil
		}
	}

	return 0, nil
}
