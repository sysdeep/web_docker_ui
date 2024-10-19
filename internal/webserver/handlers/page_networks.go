package handlers

import (
	"context"
	"net/http"
	"sort"

	"github.com/docker/docker/api/types"
	"github.com/labstack/echo/v4"
)

// models
type networkListModel struct {
	Name    string // Name is the requested name of the network
	ID      string // ID uniquely identifies a network on a single machine
	Created string // Created is the time the network created
	// Scope      string                         // Scope describes the level at which the network exists (e.g. `swarm` for cluster-wide or `local` for machine level)
	Driver string // Driver is the Driver name used to create the network (e.g. `bridge`, `overlay`)
	// EnableIPv6 bool                           // EnableIPv6 represents whether to enable IPv6
	// IPAM       network.IPAM                   // IPAM is the network's IP Address Management
	// Internal   bool                           // Internal represents if the network is used internal only
	// Attachable bool                           // Attachable represents if the global scope is manually attachable by regular containers from workers in swarm mode.
	// Ingress    bool                           // Ingress indicates the network is providing the routing-mesh for the swarm cluster.
	// ConfigFrom network.ConfigReference        // ConfigFrom specifies the source which will provide the configuration for this network.
	// ConfigOnly bool                           // ConfigOnly networks are place-holder networks for network configurations to be used by other networks. ConfigOnly networks cannot be used directly to run containers or services.
	// Containers map[string]EndpointResource    // Containers contains endpoints belonging to the network
	// Options    map[string]string              // Options holds the network specific options to use for when creating the network
	// Labels     map[string]string              // Labels holds metadata specific to the network being created
	// Peers      []network.PeerInfo             `json:",omitempty"` // List of peer nodes for an overlay network
	// Services   map[string]network.ServiceInfo `json:",omitempty"`
}
type networksPageModel struct {
	Networks []networkListModel
	Total    int
}

// handler
func (h *Handlers) NetworksPage(c echo.Context) error {
	networks_data, err := h.docker_client.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var networks []networkListModel
	for _, v := range networks_data {
		networks = append(networks, make_nerwork_list_model(v))
	}

	sort.SliceStable(networks, func(i, j int) bool {
		return networks[i].Name < networks[j].Name
	})

	response := networksPageModel{
		Networks: networks,
		Total:    len(networks),
	}

	return c.Render(http.StatusOK, "networks.html", response)
}

func make_nerwork_list_model(model types.NetworkResource) networkListModel {
	// utils.PrintAsJson(model)
	return networkListModel{
		ID:      model.ID,
		Name:    model.Name,
		Created: model.Created.String(),
		Driver:  model.Driver,
	}
}
