package services

import (
	"context"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type ContainersService struct {
	docker_client *client.Client
}

func NewContainersService(docker_client *client.Client) *ContainersService {
	return &ContainersService{docker_client: docker_client}
}

func (cs *ContainersService) GetAll() ([]ContainerListModel, error) {
	var result []ContainerListModel
	raw_containers, err := cs.docker_client.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return result, err
	}

	for _, c := range raw_containers {
		result = append(result, make_container_list_model(c))
	}

	return result, nil
}

// models ---------------------------------------------------------------------
type ContainerListModel struct {
	ID          string
	Name        string
	Image       string
	State       string
	CreatedStr  string
	IPAddresses []string
	Ports       []string
}

func make_container_list_model(c types.Container) ContainerListModel {
	// fmt.Println("------------------------------")
	// fmt.Printf("%+v\n", c)

	// utils.PrintAsJson(c)
	// time format
	unixTimeUTC := time.Unix(c.Created, 0) //gives unix time stamp in utc

	unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339) // converts utc time to RFC3339 format

	// fmt.Println("unix time stamp in UTC :--->",unixTimeUTC)
	// fmt.Println("unix time stamp in unitTimeInRFC3339 format :->",unitTimeInRFC3339)

	return ContainerListModel{
		ID:          c.ID,
		Name:        c.Names[0],
		Image:       c.Image,
		State:       c.State,
		CreatedStr:  unitTimeInRFC3339,
		IPAddresses: make([]string, 0), // TODO
		Ports:       make([]string, 0), // TODO
	}
}

/*
	"Ports": [
        {
            "IP": "127.0.0.1",
            "PrivatePort": 5000,
            "PublicPort": 5000,
            "Type": "tcp"
        },
        {
            "IP": "172.28.1.1",
            "PrivatePort": 5000,
            "PublicPort": 5000,
            "Type": "tcp"
        }
    ]
*/
