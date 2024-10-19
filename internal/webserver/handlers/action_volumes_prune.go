package handlers

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/types/filters"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) ActionVolumesPrune(c echo.Context) error {
	h.logger.Info("Prune volumes action")
	result, err := h.docker_client.VolumesPrune(context.Background(), filters.Args{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, volume_report := range result.VolumesDeleted {
		h.logger.Info("\t" + volume_report)
	}

	return c.Redirect(http.StatusFound, "/volumes")
}

// type volumesPruneReport struct {
// 	VolumesDeleted []string
// 	SpaceReclaimed uint64
// }
