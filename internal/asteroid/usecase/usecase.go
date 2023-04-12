package usecase

import (
	"context"
	"github.com/22Fariz22/asteroid-warning/internal/asteroid"
	"github.com/22Fariz22/asteroid-warning/internal/entity"
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
	"sort"
	"time"
)

type AsteroidUseCase struct {
	asteroidRepo   asteroid.AsteroidRepository
	asteroidWebAPI asteroid.NasaWebAPI
}

func NewAsteroidUseCase(asteroidRepo asteroid.AsteroidRepository, webAPI asteroid.NasaWebAPI) *AsteroidUseCase {
	return &AsteroidUseCase{
		asteroidRepo:   asteroidRepo,
		asteroidWebAPI: webAPI}
}

var datesTimeTime []time.Time

// GetNextDateUC return counts
func (a *AsteroidUseCase) GetNextDateUC(ctx context.Context, l logger.Interface, dates []string) (int64, error) {
	l.Info("GetNextDateUC.")

	//перевести даты в time.Time и отсортировать даты
	for i, _ := range dates {
		date := dates[i]
		mydate, _ := time.Parse("2006-01-02", date)
		datesTimeTime = append(datesTimeTime, mydate)
	}
	sort.Slice(datesTimeTime, func(i, j int) bool {
		return datesTimeTime[i].Before(datesTimeTime[j])
	})

	neoL, err := a.asteroidWebAPI.GetNextDateWebAPI(ctx, l, datesTimeTime)
	if err != nil {
		l.Error("err in GetNextDateUC: ", err)
		return 0, err
	}

	return neoL, nil
}

func (a *AsteroidUseCase) SaveAsteroidsUC(l logger.Interface, asteroids *entity.NeoCounts) error {
	l.Info("SaveAsteroidsUC.")
	return a.asteroidRepo.SaveAsteroidsRepo(l, asteroids)

}
