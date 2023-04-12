package usecase

import (
	"github.com/22Fariz22/asteroid-warning/internal/asteroid"
	"github.com/22Fariz22/asteroid-warning/internal/entity"
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
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

func (a *AsteroidUseCase) GetNextDateUC(l logger.Interface, dates []string) (int, error) {
	l.Info("GetNextDateUC.")
	return a.asteroidWebAPI.GetNextDateWebAPI(l, dates)
}

func (a *AsteroidUseCase) SaveAsteroidsUC(l logger.Interface, asteroids entity.NeoCounts) error {
	l.Info("SaveAsteroidsUC.")
	return a.asteroidRepo.SaveAsteroidsRepo(l, asteroids)

}

//here will some methods for usecase
