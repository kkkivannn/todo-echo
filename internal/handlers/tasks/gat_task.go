package tasks

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handlers) getTask(c echo.Context) error {
	ctx := c.Request().Context()

	paramID := c.Param("id")

	if paramID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "path param paramID can't be empty",
		})
	}

	ID, err := strconv.Atoi(paramID)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	task, err := h.ts.GetTask(ctx, ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "Not found Task",
		})
	}

	return c.JSON(http.StatusOK, task)
}
