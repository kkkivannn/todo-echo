package tasks

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handlers) getTasks(c echo.Context) error {
	ctx := c.Request().Context()

	tasks, err := h.ts.GetTasks(ctx)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, tasks)
}
