package webAPI

import (
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
)

type WebAPI struct {
}

func NewWebAPI() *WebAPI {
	return &WebAPI{}
}

func (w *WebAPI) GetNextDateWebAPI(l logger.Interface, dates []string) (int, error) {
	l.Info("GetNextDateWebAPI.")

	//добавить составить даты из апи nasa
	//вернуть пользователю количество астероидов по ближащей дате

	return 0, nil
}
