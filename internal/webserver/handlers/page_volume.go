package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types/volume"
	"github.com/labstack/echo/v4"
)

// TODO: used by containers

// models
type volumeModel struct {
	Name       string
	CreatedAt  string
	Driver     string
	Mountpoint string
}

type volumePageModel struct {
	Volume volumeModel
}

// handler
func (h *Handlers) VolumePage(c echo.Context) error {
	name := c.Param("name")
	volume_data, err := h.docker_client.VolumeInspect(context.Background(), name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	fmt.Printf("%+v\n", volume_data)

	response := volumePageModel{
		Volume: make_volume_model(&volume_data),
	}

	return c.Render(http.StatusOK, "volume.html", response)
}

func make_volume_model(data *volume.Volume) volumeModel {
	return volumeModel{
		Name:       data.Name,
		Driver:     data.Driver,
		CreatedAt:  data.CreatedAt,
		Mountpoint: data.Mountpoint,
	}
}
