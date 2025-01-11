package tasks

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type RequestUpdateTask struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	StatusID int    `json:"statusID"`
}

func (h *Handlers) updateTask(c echo.Context) error {
	ctx := c.Request().Context()

	r := new(RequestUpdateTask)

	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "title or body not valid",
		})
	}

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

	task, err := h.ts.EditTask(ctx, ID, r.Title, r.Body, r.StatusID)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, task)
}
