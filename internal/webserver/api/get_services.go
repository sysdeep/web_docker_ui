package api

import (
	"context"
	"hdu/internal/utils"
	"net/http"
	"sort"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/labstack/echo/v4"
)

// handler
func (h *Api) GetServices(c echo.Context) error {
	services_list, err := h.docker_client.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	response := newServicesPageModel(services_list)

	return c.JSON(http.StatusOK, response)
}

// models ---------------------------------------------------------------------
type serviceSpec struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
}

// Service represents a service.
type serviceListModel struct {
	ID        string      `json:"id"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	Spec      serviceSpec `json:"spec"`
	// PreviousSpec *ServiceSpec  `json:",omitempty"`
	// Endpoint     Endpoint      `json:",omitempty"`
	// UpdateStatus *UpdateStatus `json:",omitempty"`
	//
	// // ServiceStatus is an optional, extra field indicating the number of
	// // desired and running tasks. It is provided primarily as a shortcut to
	// // calculating these values client-side, which otherwise would require
	// // listing all tasks for a service, an operation that could be
	// // computation and network expensive.
	// ServiceStatus *ServiceStatus `json:",omitempty"`
	//
	// // JobStatus is the status of a Service which is in one of ReplicatedJob or
	// // GlobalJob modes. It is absent on Replicated and Global services.
	// JobStatus *JobStatus `json:",omitempty"`
}

func newServiceListModel(model swarm.Service) serviceListModel {
	utils.PrintAsJson(model)

	created_str := model.CreatedAt.Format(time.RFC3339) // converts utc time to RFC3339 format
	updated_str := model.UpdatedAt.Format(time.RFC3339) // converts utc time to RFC3339 format

	spec := serviceSpec{
		Name:   model.Spec.Name,
		Labels: model.Spec.Labels,
	}
	return serviceListModel{
		ID:        model.ID,
		CreatedAt: created_str,
		UpdatedAt: updated_str,
		Spec:      spec,
	}
}

// page model -----------------------------------------------------------------
type servicesPageModel struct {
	Services []serviceListModel `json:"services"`
	Total    int                `json:"total"`
}

func newServicesPageModel(models []swarm.Service) servicesPageModel {
	services := []serviceListModel{}

	for _, model := range models {
		services = append(services, newServiceListModel(model))
	}

	sort.SliceStable(services, func(i, j int) bool {
		return services[i].ID < services[j].ID
	})

	return servicesPageModel{
		Services: services,
		Total:    len(services),
	}
}
