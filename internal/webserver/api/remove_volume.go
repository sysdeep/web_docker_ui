package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Api) RemoveVolume(c echo.Context) error {
	id := c.Param("name")
	h.logger.Info("remove volume: " + id)

	force := false

	err := h.docker_client.VolumeRemove(context.Background(), id, force)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, stdResponse{true})
}
