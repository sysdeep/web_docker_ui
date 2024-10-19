package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) ActionConfigRemove(c echo.Context) error {
	id := c.Param("id")
	h.logger.Info("remove config: " + id)

	err := h.docker_client.ConfigRemove(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.Redirect(http.StatusFound, "/configs")
}
