package app

import (
	"github.com/labstack/echo/v4"
	"todo_echo/internal/config"
	"todo_echo/internal/handlers"
	"todo_echo/internal/http"
	storage "todo_echo/internal/storage/sqlite"
)

type App struct {
	cfg *config.Config
	Srv *http.Server
}

func New(cfg *config.Config) *App {
	e := echo.New()

	srv := http.New(e, cfg.Host, cfg.Port)

	s := storage.New(cfg.DBString)

	h := handlers.New(e, s)

	h.SetupRoutes()

	return &App{
		cfg: cfg,
		Srv: srv,
	}
}
