package service_model

import (
	"hdu/internal/utils"
	"time"

	"github.com/docker/docker/api/types/swarm"
)

func NewServiceListModel(model swarm.Service) ServiceModel {
	utils.PrintAsJson(model)

	created_str := model.CreatedAt.Format(time.RFC3339) // converts utc time to RFC3339 format
	updated_str := model.UpdatedAt.Format(time.RFC3339) // converts utc time to RFC3339 format

	var global_mode *GlobalService
	if model.Spec.Mode.Global != nil {
		global_mode = &GlobalService{}
	}

	var replicated_mode *ReplicatedService
	if model.Spec.Mode.Replicated != nil {
		var replicas_count int
		if model.Spec.Mode.Replicated.Replicas == nil {
			replicas_count = 0
		} else {
			replicas_count = int(*model.Spec.Mode.Replicated.Replicas)
		}
		replicated_mode = &ReplicatedService{
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
	return ServiceModel{
		ID:        model.ID,
		Name:      model.Spec.Name,
		CreatedAt: created_str,
		UpdatedAt: updated_str,
		Image:     image,
		Mode: ServiceMode{
			Replicated: replicated_mode,
			Global:     global_mode,
		},

		// Spec:      spec,
	}
}
