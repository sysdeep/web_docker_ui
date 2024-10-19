package handlers

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/labstack/echo/v4"
)

// models
type networModel struct {
	Name    string // Name is the requested name of the network
	ID      string // ID uniquely identifies a network on a single machine
	Created string // Created is the time the network created
	Scope   string // Scope describes the level at which the network exists (e.g. `swarm` for cluster-wide or `local` for machine level)
	Driver  string // Driver is the Driver name used to create the network (e.g. `bridge`, `overlay`)
	// EnableIPv6 bool                           // EnableIPv6 represents whether to enable IPv6
	// IPAM       network.IPAM                   // IPAM is the network's IP Address Management
	Internal   bool // Internal represents if the network is used internal only
	Attachable bool // Attachable represents if the global scope is manually attachable by regular containers from workers in swarm mode.
	Ingress    bool // Ingress indicates the network is providing the routing-mesh for the swarm cluster.
	// ConfigFrom network.ConfigReference        // ConfigFrom specifies the source which will provide the configuration for this network.
	// ConfigOnly bool                           // ConfigOnly networks are place-holder networks for network configurations to be used by other networks. ConfigOnly networks cannot be used directly to run containers or services.
	Containers []networkContainerModel // Containers contains endpoints belonging to the network
	// Options    map[string]string              // Options holds the network specific options to use for when creating the network
	// Labels     map[string]string              // Labels holds metadata specific to the network being created
	// Peers      []network.PeerInfo             `json:",omitempty"` // List of peer nodes for an overlay network
	// Services   map[string]network.ServiceInfo `json:",omitempty"`
}

type networkContainerModel struct {
	ID          string
	Name        string
	EndpointID  string
	MacAddress  string
	IPv4Address string
	IPv6Address string
}

type networkPageModel struct {
	Network networModel
}

// handler
func (h *Handlers) NetworkPage(c echo.Context) error {
	id := c.Param("id")
	network_data, err := h.docker_client.NetworkInspect(context.Background(), id, types.NetworkInspectOptions{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	response := networkPageModel{
		Network: make_nerwork_model(network_data),
	}

	return c.Render(http.StatusOK, "network.html", response)
}

func make_nerwork_model(model types.NetworkResource) networModel {
	// utils.PrintAsJson(model)

	// make containers
	var containers []networkContainerModel
	if model.Containers != nil {
		for key, cmodel := range model.Containers {
			container := networkContainerModel{
				ID:          key,
				Name:        cmodel.Name,
				EndpointID:  cmodel.EndpointID,
				MacAddress:  cmodel.MacAddress,
				IPv4Address: cmodel.IPv4Address,
				IPv6Address: cmodel.IPv6Address,
			}

			containers = append(containers, container)
		}
	}

	return networModel{
		ID:         model.ID,
		Name:       model.Name,
		Created:    model.Created.String(),
		Driver:     model.Driver,
		Scope:      model.Scope,
		Internal:   model.Internal,
		Attachable: model.Attachable,
		Ingress:    model.Ingress,
		Containers: containers,
	}
}
