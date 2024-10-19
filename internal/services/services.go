package services

import "github.com/docker/docker/client"

type Services struct {
	Containers *ContainersService
	Images     *ImagesService
	Volumes    *VolumesService
}

func NewServices(docker_client *client.Client) *Services {

	return &Services{
		Containers: NewContainersService(docker_client),
		Images:     NewImagesService(docker_client),
		Volumes:    NewVolumesService(docker_client),
	}
}
