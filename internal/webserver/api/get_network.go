package api

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/types/network"
	"github.com/labstack/echo/v4"
)

// models
type networkModel struct {
	Name    string `json:"name"`    // Name is the requested name of the network
	ID      string `json:"id"`      // ID uniquely identifies a network on a single machine
	Created string `json:"created"` // Created is the time the network created
	Scope   string `json:"scope"`   // Scope describes the level at which the network exists (e.g. `swarm` for cluster-wide or `local` for machine level)
	Driver  string `json:"driver"`  // Driver is the Driver name used to create the network (e.g. `bridge`, `overlay`)
	// EnableIPv6 bool                           // EnableIPv6 represents whether to enable IPv6
	// IPAM       network.IPAM                   // IPAM is the network's IP Address Management
	Internal   bool `json:"internal"`   // Internal represents if the network is used internal only
	Attachable bool `json:"attachable"` // Attachable represents if the global scope is manually attachable by regular containers from workers in swarm mode.
	Ingress    bool `json:"ingress"`    // Ingress indicates the network is providing the routing-mesh for the swarm cluster.
	// ConfigFrom network.ConfigReference        // ConfigFrom specifies the source which will provide the configuration for this network.
	// ConfigOnly bool                           // ConfigOnly networks are place-holder networks for network configurations to be used by other networks. ConfigOnly networks cannot be used directly to run containers or services.
	// Options    map[string]string              // Options holds the network specific options to use for when creating the network
	// Labels     map[string]string              // Labels holds metadata specific to the network being created
	// Peers      []network.PeerInfo             `json:",omitempty"` // List of peer nodes for an overlay network
	// Services   map[string]network.ServiceInfo `json:",omitempty"`
}

type networkContainerModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	EndpointID  string `json:"endpoint_id"`
	MacAddress  string `json:"mac_address"`
	IPv4Address string `json:"ip_v4_address"`
	IPv6Address string `json:"ip_v6_address"`
}

type networkPageModel struct {
	Network    networkModel            `json:"network"`
	Containers []networkContainerModel `json:"containers"` // Containers contains endpoints belonging to the network
}

// handler
func (h *Api) GetNetwork(c echo.Context) error {
	id := c.Param("id")
	network_data, err := h.docker_client.NetworkInspect(context.Background(), id, network.InspectOptions{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	containers := []networkContainerModel{}
	if network_data.Containers != nil {
		for key, cmodel := range network_data.Containers {
			container := make_network_container_model(key, cmodel)
			containers = append(containers, container)
		}
	}

	response := networkPageModel{
		Network:    make_nerwork_model(network_data),
		Containers: containers,
	}

	return c.JSON(http.StatusOK, response)
}

// Containers map[string]EndpointResource // Containers contains endpoints belonging to the network
func make_network_container_model(id string, model network.EndpointResource) networkContainerModel {
	return networkContainerModel{
		ID:          id,
		Name:        model.Name,
		EndpointID:  model.EndpointID,
		MacAddress:  model.MacAddress,
		IPv4Address: model.IPv4Address,
		IPv6Address: model.IPv6Address,
	}
}

func make_nerwork_model(model network.Inspect) networkModel {
	// utils.PrintAsJson(model)

	return networkModel{
		ID:         model.ID,
		Name:       model.Name,
		Created:    model.Created.String(),
		Driver:     model.Driver,
		Scope:      model.Scope,
		Internal:   model.Internal,
		Attachable: model.Attachable,
		Ingress:    model.Ingress,
	}
}
