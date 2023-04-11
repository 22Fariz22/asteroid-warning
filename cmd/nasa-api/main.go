package main

import (
	"github.com/22Fariz22/asteroid-warning/internal/app"
	"github.com/22Fariz22/asteroid-warning/internal/config"
)

func main() {
	cfg := config.NewConfig()

	app := app.NewApp(cfg)
	app.Run()
}

//запуск приложения
//go run cmd/nasa-api/main.go -d="postgres://postgres:55555@127.0.0.1:5432/asteroids" -a="localhost:8080"
