package docker_adapter

import (
	"hdu/internal/webserver/api/models"

	"github.com/docker/docker/api/types/swarm"
)

func NewSecretModel(model swarm.Secret) models.Secret {
	// utils.PrintAsJson(model)
	return models.Secret{
		ID:      model.ID,
		Name:    model.Spec.Name,
		Created: model.CreatedAt.String(),
		Updated: model.UpdatedAt.String(),
	}
}
