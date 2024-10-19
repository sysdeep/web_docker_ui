package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) ActionVolumeRemove(c echo.Context) error {
	name := c.Param("name")
	h.logger.Info("remove volume: " + name)

	force := false

	err := h.docker_client.VolumeRemove(context.Background(), name, force)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.Redirect(http.StatusFound, "/volumes")
}
