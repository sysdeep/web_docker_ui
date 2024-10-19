package api

import (
	"hdu/internal/logger"
	"hdu/internal/registry_client"
	"hdu/internal/services"

	"github.com/docker/docker/client"
)

type Api struct {
	docker_client   *client.Client
	logger          *logger.Logger
	services        *services.Services
	registry_client *registry_client.RegistryClient
}

func NewApi(docker_client *client.Client, registry_client *registry_client.RegistryClient, services *services.Services, logger *logger.Logger) *Api {

	return &Api{
		docker_client:   docker_client,
		logger:          logger,
		services:        services,
		registry_client: registry_client,
	}
}
