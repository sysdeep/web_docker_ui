package api

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/types/image"
	"github.com/labstack/echo/v4"
)

func (h *Api) RemoveImage(c echo.Context) error {
	id := c.Param("id")
	h.logger.Info("remove image: " + id)

	options := image.RemoveOptions{
		Force:         true,
		PruneChildren: false,
	}
	_, err := h.docker_client.ImageRemove(context.Background(), id, options)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, stdResponse{true})
}
