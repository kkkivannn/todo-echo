package tasks

import (
	"context"
	"github.com/labstack/echo/v4"
	"todo_echo/internal/domain/models"
	"todo_echo/internal/services"
)

type TaskService interface {
	AddTask(ctx context.Context, title, body string) (int, error)
	GetTask(ctx context.Context, taskID int) (models.Task, error)
	GetTasks(ctx context.Context) ([]models.Task, error)
	RemoveTask(ctx context.Context, taskID int) error
	EditTask(ctx context.Context, taskID int, title string, body string, statusID int) (models.Task, error)
}

type Handlers struct {
	e  *echo.Echo
	ts *services.Tasks
}

func New(e *echo.Echo, ts *services.Tasks) *Handlers {
	return &Handlers{e: e, ts: ts}
}

func (h *Handlers) InitRoutes() {
	task := h.e.Group("/task")
	task.POST("", h.addTask)
	task.DELETE("/:id", h.deleteTask)
	task.GET("/:id", h.getTask)
	task.PATCH("/:id", h.updateTask)
	h.e.GET("/tasks", h.getTasks)
}
