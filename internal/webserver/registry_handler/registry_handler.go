package registry_handler

import (
	"hdu/internal/registry_client"

	"github.com/docker/docker/client"
)

type RegistryHandler struct {
	docker_client   *client.Client
	registry_client *registry_client.RegistryClient
}

func NewRegistryHandler(docker_client *client.Client, registry_client *registry_client.RegistryClient) *RegistryHandler {
	return &RegistryHandler{
		docker_client:   docker_client,
		registry_client: registry_client,
	}
}

type nullResponse struct {
}
