package handlers

import (
	"hdu/internal/logger"
	"hdu/internal/services"

	"github.com/docker/docker/client"
)

type Handlers struct {
	docker_client *client.Client
	logger        *logger.Logger
	services      *services.Services
}

func NewHandlers(docker_client *client.Client, services *services.Services, logger *logger.Logger) *Handlers {

	return &Handlers{
		docker_client: docker_client,
		logger:        logger,
		services:      services,
	}
}
