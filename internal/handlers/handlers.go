package handlers

import (
	"log"
	db "todo_echo/internal/storage/sqlite"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	srv *echo.Echo
	db  *db.TaskStorage
}

func New(srv *echo.Echo, db *db.TaskStorage) *Handler {
	return &Handler{srv: srv, db: db}
}

func (h *Handler) SetupRoutes() {
	h.srv.GET("/health", h.HealthCheck)
}

func (h *Handler) HealthCheck(e echo.Context) error {

	err := h.db.HealthCheck()
	if err != nil {
		log.Fatal(err)
		return e.JSON(500, "Internal Server Error")
	}

	log.Println("Health check passed")
	return e.JSON(200, map[string]interface{}{
		"status": "ok",
	})
}
