package api

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/labstack/echo/v4"
)

// TODO: used by containers

// models
type volumeModel struct {
	Name       string `json:"name"`
	CreatedAt  string `json:"created"`
	Driver     string `json:"driver"`
	Mountpoint string `json:"mount_point"`
}

type volumePageModel struct {
	Volume     volumeModel          `json:"volume"`
	Containers []containerListModel `json:"containers"`
}

// handler
func (h *Api) GetVolume(c echo.Context) error {
	name := c.Param("name")

	// volume data
	volume_data, err := h.docker_client.VolumeInspect(context.Background(), name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// containers
	// TODO: получить список контейнеров в которых используется данный волум
	// https://localhost:9443/api/endpoints/2/docker/containers/json?all=1&filters=%7B%22volume%22:%5B%22portainer_data%22%5D%7D
	// fmt.Printf("%+v\n", volume_data)

	containers_filter := filters.NewArgs(filters.KeyValuePair{Key: "volume", Value: name})
	// Получение списка запуцщенных контейнеров(docker ps)
	raw_containers, err := h.docker_client.ContainerList(context.Background(), container.ListOptions{All: true, Filters: containers_filter})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	containers := []containerListModel{}
	for _, c := range raw_containers {
		containers = append(containers, convert_container(c))
	}

	// response
	response := volumePageModel{
		Volume:     make_volume_model(&volume_data),
		Containers: containers,
	}

	return c.JSON(http.StatusOK, response)
}

func make_volume_model(data *volume.Volume) volumeModel {
	return volumeModel{
		Name:       data.Name,
		Driver:     data.Driver,
		CreatedAt:  data.CreatedAt,
		Mountpoint: data.Mountpoint,
	}
}
