package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Api) RemoveNetwork(c echo.Context) error {
	id := c.Param("id")
	h.logger.Info("remove network: " + id)

	err := h.docker_client.NetworkRemove(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, stdResponse{true})
}
