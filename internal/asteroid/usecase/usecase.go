package usecase

import (
	"github.com/22Fariz22/asteroid-warning/internal/asteroid"
)

type AsteroidUseCase struct {
	asteroidRepo asteroid.AsteroidRepository
}

func NewAsteroidUseCase(asteroidRepo asteroid.AsteroidRepository) *AsteroidUseCase {
	return &AsteroidUseCase{asteroidRepo: asteroidRepo}
}

//here will some methods for usecase
