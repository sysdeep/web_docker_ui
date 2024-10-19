package api

import (
	"hdu/internal/services"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
)

// models
type volumesPageModel struct {
	Volumes []ApiVolumeListModel `json:"volumes"`
	Total   int                  `json:"total"`
}

// handler
func (h *Api) GetVolumes(c echo.Context) error {

	volumes, err := h.services.Volumes.GetAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	sort.SliceStable(volumes, func(i, j int) bool {
		return volumes[i].Name < volumes[j].Name
	})

	api_volumes := []ApiVolumeListModel{}
	for _, vol := range volumes {
		api_volumes = append(api_volumes, toApiVolumeListModel(&vol))
	}

	response := volumesPageModel{
		Volumes: api_volumes,
		Total:   len(api_volumes),
	}

	return c.JSON(http.StatusOK, response)
}

type ApiVolumeListModel struct {
	Name       string `json:"name"`
	CreatedAt  string `json:"created"`
	Driver     string `json:"driver"`
	Mountpoint string `json:"mount_point"`
	StackName  string `json:"stack_name"`
	Used       bool   `json:"used"`
}

func toApiVolumeListModel(model *services.VolumeListModel) ApiVolumeListModel {
	return ApiVolumeListModel{
		Name:       model.Name,
		CreatedAt:  model.CreatedAt,
		Driver:     model.Driver,
		Mountpoint: model.Mountpoint,
		StackName:  model.StackName,
		Used:       model.Used,
	}
}
