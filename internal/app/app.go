package app

import (
	"todo_echo/internal/config"
	"todo_echo/internal/handlers"
	"todo_echo/internal/http"
	"todo_echo/internal/services"
	storage "todo_echo/internal/storage/sqlite"

	"github.com/labstack/echo/v4"
)

type App struct {
	cfg *config.Config
	Srv *http.Server
}

func New(cfg *config.Config) *App {
	e := echo.New()

	srv := http.New(e, cfg.Host, cfg.Port)

	s, err := storage.New(cfg.DBString)
	if err != nil {
		panic(err)
	}

	taskService := services.New(s)

	h := handlers.New(e, taskService)

	h.SetupRoutes()

	return &App{
		cfg: cfg,
		Srv: srv,
	}
}
