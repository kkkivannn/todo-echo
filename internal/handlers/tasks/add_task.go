package tasks

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RequestAddTask struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (h *Handlers) addTask(c echo.Context) error {

	r := new(RequestAddTask)

	ctx := c.Request().Context()

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "title or body not valid",
		})
	}

	id, err := h.ts.AddTask(ctx, r.Title, r.Body)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}
