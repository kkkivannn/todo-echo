package tasks

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handlers) deleteTask(c echo.Context) error {
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

	err = h.ts.RemoveTask(ctx, ID)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
