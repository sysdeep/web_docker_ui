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

// handler --------------------------------------------------------------------
func (h *Api) GetServices(c echo.Context) error {
	services_list, err := h.docker_client.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	response := newServicesPageModel(services_list)

	return c.JSON(http.StatusOK, response)
}

// models ---------------------------------------------------------------------
// type serviceSpec struct {
// 	Name   string            `json:"name"`
// 	Labels map[string]string `json:"labels"`
// }

type serviceMode struct {
	Replicated *replicatedService `json:"replicated"`
	Global     *globalService     `json:"global"`
	// ReplicatedJob *ReplicatedJob     `json:",omitempty"`
	// GlobalJob     *GlobalJob         `json:",omitempty"`
}

// ReplicatedService is a kind of ServiceMode.
type replicatedService struct {
	Replicas int `json:"replicas"`
}

// GlobalService is a kind of ServiceMode.
type globalService struct{}

// Service represents a service.
type serviceListModel struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Mode      serviceMode `json:"mode"`
	Image     string      `json:"image"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	// Spec      serviceSpec `json:"spec"`
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

	var global_mode *globalService
	if model.Spec.Mode.Global != nil {
		global_mode = &globalService{}
	}

	var replicated_mode *replicatedService
	if model.Spec.Mode.Replicated != nil {
		var replicas_count int
		if model.Spec.Mode.Replicated.Replicas == nil {
			replicas_count = 0
		} else {
			replicas_count = int(*model.Spec.Mode.Replicated.Replicas)
		}
		replicated_mode = &replicatedService{
			Replicas: replicas_count,
		}
	}

	var image string
	if model.Spec.TaskTemplate.ContainerSpec != nil {
		image = model.Spec.TaskTemplate.ContainerSpec.Image
	}

	// spec := serviceSpec{
	// 	Name:   model.Spec.Name,
	// 	Labels: model.Spec.Labels,
	// }
	return serviceListModel{
		ID:        model.ID,
		Name:      model.Spec.Name,
		CreatedAt: created_str,
		UpdatedAt: updated_str,
		Image:     image,
		Mode: serviceMode{
			Replicated: replicated_mode,
			Global:     global_mode,
		},

		// Spec:      spec,
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
