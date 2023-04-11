package app

import (
	"context"
	"fmt"
	"github.com/22Fariz22/asteroid-warning/internal/asteroid"
	asteroidRouter "github.com/22Fariz22/asteroid-warning/internal/asteroid/delivery/http"
	asteroidRepo "github.com/22Fariz22/asteroid-warning/internal/asteroid/repository/postgres"
	"github.com/22Fariz22/asteroid-warning/internal/asteroid/usecase"
	"github.com/22Fariz22/asteroid-warning/internal/config"
	"github.com/22Fariz22/asteroid-warning/pkg/logger"
	"github.com/22Fariz22/asteroid-warning/pkg/postgres"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	cfg        *config.Config
	httpServer *http.Server
	asteroidUC asteroid.UseCase
}

func NewApp(cfg *config.Config) *App {
	// Repository
	db, err := postgres.New(cfg.DatabaseURI, postgres.MaxPoolSize(2))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	asteroidRepo := asteroidRepo.NewAsteroidRepository(db)

	return &App{
		cfg:        cfg,
		httpServer: nil,
		asteroidUC: usecase.NewAsteroidUseCase(asteroidRepo),
	}
}

func (a *App) Run() {
	l := logger.New("debug")

	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	api := router.Group("/")

	asteroidRouter.RegisterHTTPEndpointsOrder(api, a.asteroidUC, l)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           a.cfg.RunAddress,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			l.Fatal("Failed to listen and serve: %+v", err)
		}
	}()

	//gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	a.httpServer.Shutdown(ctx)

}
