package usecase

import (
	"github.com/22Fariz22/asteroid-warning/internal/asteroid"
	"github.com/22Fariz22/asteroid-warning/internal/entity"
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
)

type AsteroidUseCase struct {
	asteroidRepo asteroid.AsteroidRepository
}

func NewAsteroidUseCase(asteroidRepo asteroid.AsteroidRepository) *AsteroidUseCase {
	return &AsteroidUseCase{asteroidRepo: asteroidRepo}
}

func (a *AsteroidUseCase) GetNextDateUC(l logger.Interface, dates []string) (*entity.Asteroid, error) {
	l.Info("GetNextDateUC.")
	return a.asteroidRepo.GetNextDateRepo(l, dates)
}

func (a *AsteroidUseCase) SaveAsteroidsUC(l logger.Interface, asteroids entity.NeoCounts) error {
	l.Info("SaveAsteroidsUC.")
	return a.asteroidRepo.SaveAsteroidsRepo(l, asteroids)

}

//here will some methods for usecase
