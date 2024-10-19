package api

import (
	"context"
	"hdu/internal/utils"
	"net/http"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/labstack/echo/v4"
)

// models
type containerListModel struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Image       string   `json:"image"`
	ImageID     string   `json:"image_id"`
	State       string   `json:"state"`
	CreatedStr  string   `json:"created"`
	IPAddresses []string `json:"ip_addresses"`
	Ports       []string //TODO
}

// type containerListNetworkModel struct {
// 	Type      string `json:"type"` // bridge
// 	IPAddress string `json:"ip_address"`
// }

type containersPageModel struct {
	Containers []containerListModel `json:"containers"`
}

// handler
func (h *Api) GetContainers(c echo.Context) error {
	// Получение списка запуцщенных контейнеров(docker ps)
	raw_containers, err := h.docker_client.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	utils.PrintAsJson(raw_containers)
	containers := []containerListModel{}
	for _, c := range raw_containers {
		containers = append(containers, convert_container(c))
	}

	response := containersPageModel{
		Containers: containers,
	}

	return c.JSON(http.StatusOK, response)
}

/*
	{
		ID:049011df54ab6f32b79adb2b36ee84456581e70ef61142c213de59338fcb9234
		Names:[/first_real_motohours_motohours_1]
		Image:comcon/node/motohours:1.0.1
		ImageID:sha256:3b564457197a55585a3666ba49eb9a5820a2edf5aa8d0e21d763d08a94c6e9ae
		Command:npm start
		Created:1690526956
		Ports:[]
		SizeRw:0
		SizeRootFs:0
		Labels:map[
			com.docker.compose.config-hash:8f0f92546a097e094f93c43d32a0fdd5f20e3426394f426b60ee77c3b002f876
			com.docker.compose.container-number:1
			com.docker.compose.oneoff:False
			com.docker.compose.project:first_real_motohours
			com.docker.compose.project.config_files:docker-compose.yml
			com.docker.compose.project.working_dir:/home/nia/Development/Comcon/comcon-scripts/comcon-scripts/docker/stacks/first_real_motohours
			com.docker.compose.service:motohours com.docker.compose.version:1.29.2
			]
		State:exited
		Status:Exited (137) 9 months ago
		HostConfig:{NetworkMode:host}
		NetworkSettings:0xc00034c090
		Mounts:[
			{
				Type:bind
				Name: Source:/home/nia/Development/Comcon/comcon-scripts/comcon-scripts/docker/stacks/first_real_motohours/opt/motohours/config.json
				Destination:/app/config.json
				Driver:
				Mode:rw
				RW:true
				Propagation:rprivate
				}
				]
			}
*/
func convert_container(c types.Container) containerListModel {
	// fmt.Println("------------------------------")
	// fmt.Printf("%+v\n", c)

	// time format
	unixTimeUTC := time.Unix(c.Created, 0) //gives unix time stamp in utc

	unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339) // converts utc time to RFC3339 format

	// fmt.Println("unix time stamp in UTC :--->",unixTimeUTC)
	// fmt.Println("unix time stamp in unitTimeInRFC3339 format :->",unitTimeInRFC3339)

	// extract ip_address
	ip_addresses := []string{}
	for _, net_config := range c.NetworkSettings.Networks {
		ip_addresses = append(ip_addresses, net_config.IPAddress)
	}

	return containerListModel{
		ID:          c.ID,
		Name:        c.Names[0],
		Image:       c.Image,
		ImageID:     c.ImageID,
		State:       c.State,
		CreatedStr:  unitTimeInRFC3339,
		IPAddresses: ip_addresses,      // TODO
		Ports:       make([]string, 0), // TODO
	}
}
