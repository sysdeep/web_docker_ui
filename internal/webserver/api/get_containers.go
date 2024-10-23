package api

import (
	"context"
	"hdu/internal/utils"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/labstack/echo/v4"
)

// models

// type containerListNetworkModel struct {
// 	Type      string `json:"type"` // bridge
// 	IPAddress string `json:"ip_address"`
// }

type containersPageModel struct {
	Containers []containerListModel `json:"containers"`
}

// handler
func (h *Api) GetContainers(c echo.Context) error {
	// Получение списка запуцщенных контейнеров(docker ps)
	raw_containers, err := h.docker_client.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	utils.PrintAsJson(raw_containers)
	containers := []containerListModel{}
	for _, c := range raw_containers {
		containers = append(containers, newContainerListModel(c))
	}

	response := containersPageModel{
		Containers: containers,
	}

	return c.JSON(http.StatusOK, response)
}
