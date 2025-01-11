package handlers

import (
	"github.com/labstack/echo/v4"
	"todo_echo/internal/handlers/tasks"
	"todo_echo/internal/services"
)

type Handlers struct {
	e  *echo.Echo
	t  *tasks.Handlers
	ts *services.Tasks
}

func New(e *echo.Echo, ts *services.Tasks) *Handlers {
	t := tasks.New(e, ts)
	return &Handlers{e: e, t: t}
}

func (h *Handlers) SetupRoutes() {
	h.t.InitRoutes()
}
