package handlers

import (
	"hdu/internal/services"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
)

// models
type volumesPageModel struct {
	Volumes []services.VolumeListModel
	Total   int
}

// handler
func (h *Handlers) VolumesPage(c echo.Context) error {

	volumes, err := h.services.Volumes.GetAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	sort.SliceStable(volumes, func(i, j int) bool {
		return volumes[i].Name < volumes[j].Name
	})

	response := volumesPageModel{
		Volumes: volumes,
		Total:   len(volumes),
	}

	return c.Render(http.StatusOK, "volumes.html", response)
}
