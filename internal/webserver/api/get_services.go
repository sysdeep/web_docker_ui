package api

import (
	"context"
	"hdu/internal/webserver/api/docker_adapter"
	"hdu/internal/webserver/api/models"
	"net/http"
	"sort"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/labstack/echo/v4"
)

// handler --------------------------------------------------------------------
func (h *Api) GetServices(c echo.Context) error {
	services_list, err := h.docker_client.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	response := newServicesPageModel(services_list)

	return c.JSON(http.StatusOK, response)
}

// page model -----------------------------------------------------------------
type servicesPageModel struct {
	Services []models.Service `json:"services"`
	Total    int              `json:"total"`
}

func newServicesPageModel(all_models []swarm.Service) servicesPageModel {
	services := []models.Service{}

	for _, model := range all_models {
		services = append(services, docker_adapter.NewServiceModel(model))
	}

	sort.SliceStable(services, func(i, j int) bool {
		return services[i].ID < services[j].ID
	})

	return servicesPageModel{
		Services: services,
		Total:    len(services),
	}
}
