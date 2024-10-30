package docker_adapter

import (
	"hdu/internal/utils"
	"hdu/internal/webserver/api/models"
	"time"

	"github.com/docker/docker/api/types/swarm"
)

func NewServiceModel(model swarm.Service) models.Service {
	utils.PrintAsJson(model)

	created_str := model.CreatedAt.Format(time.RFC3339) // converts utc time to RFC3339 format
	updated_str := model.UpdatedAt.Format(time.RFC3339) // converts utc time to RFC3339 format

	service_mode := newServiceMode(model.Spec.Mode)

	var image string
	if model.Spec.TaskTemplate.ContainerSpec != nil {
		image = model.Spec.TaskTemplate.ContainerSpec.Image
	}

	// spec := serviceSpec{
	// 	Name:   model.Spec.Name,
	// 	Labels: model.Spec.Labels,
	// }
	return models.Service{
		ID:        model.ID,
		Name:      model.Spec.Name,
		CreatedAt: created_str,
		UpdatedAt: updated_str,
		Image:     image,
		Mode:      service_mode,

		Spec: newServiceSpec(model.Spec),
	}
}

func newServiceSpec(spec swarm.ServiceSpec) models.ServiceSpec {
	return models.ServiceSpec{
		Name:         spec.Name,
		TaskTemplate: newServiceTaskSpec(spec.TaskTemplate),
		Mode:         newServiceMode(spec.Mode),
		EndpointSpec: newEndpointSpec(spec.EndpointSpec),
	}
}

func newServiceTaskSpec(spec swarm.TaskSpec) models.ServiceTaskSpec {
	return models.ServiceTaskSpec{
		ContainerSpec: newServiceContainerSpec(spec.ContainerSpec),
	}
}

func newServiceContainerSpec(spec *swarm.ContainerSpec) *models.ServiceContainerSpec {
	if spec == nil {
		return nil
	}

	return &models.ServiceContainerSpec{
		Image: spec.Image,
	}
}

func newServiceMode(spec swarm.ServiceMode) models.ServiceMode {
	var global_mode *models.GlobalService
	if spec.Global != nil {
		global_mode = &models.GlobalService{}
	}

	var replicated_mode *models.ReplicatedService
	if spec.Replicated != nil {
		var replicas_count int
		if spec.Replicated.Replicas == nil {
			replicas_count = 0
		} else {
			replicas_count = int(*spec.Replicated.Replicas)
		}
		replicated_mode = &models.ReplicatedService{
			Replicas: replicas_count,
		}
	}

	var replicated_job *models.ReplicatedJob
	if spec.Global != nil {
		replicated_job = &models.ReplicatedJob{}
	}

	var global_job *models.GlobalJob
	if spec.Global != nil {
		global_job = &models.GlobalJob{}
	}

	return models.ServiceMode{
		Global:        global_mode,
		Replicated:    replicated_mode,
		ReplicatedJob: replicated_job,
		GlobalJob:     global_job,
	}
}

func newEndpointSpec(spec *swarm.EndpointSpec) *models.EndpointSpec {
	if spec == nil {
		return nil
	}

	ports_config := []models.PortConfig{}
	for _, pc := range spec.Ports {
		npc := models.PortConfig{
			Name:          pc.Name,
			Protocol:      string(pc.Protocol),
			TargetPort:    pc.TargetPort,
			PublishedPort: pc.PublishedPort,
			PublishMode:   string(pc.PublishMode),
		}

		ports_config = append(ports_config, npc)
	}

	return &models.EndpointSpec{
		Mode:  string(spec.Mode),
		Ports: ports_config,
	}
}
