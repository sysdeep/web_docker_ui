package handlers

import (
	"context"
	"net/http"
	"sort"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/labstack/echo/v4"
)

// models
type configListModel struct {
	ID      string
	Created string
	Updated string
	Name    string
}

type configsPageModel struct {
	Configs []configListModel
	Total   int
}

// handler
func (h *Handlers) ConfigsPage(c echo.Context) error {
	configs_data, err := h.docker_client.ConfigList(context.Background(), types.ConfigListOptions{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var configs []configListModel
	for _, v := range configs_data {
		configs = append(configs, make_config_list_model(v))
	}

	sort.SliceStable(configs, func(i, j int) bool {
		return configs[i].Name < configs[j].Name
	})

	response := configsPageModel{
		Configs: configs,
		Total:   len(configs),
	}

	return c.Render(http.StatusOK, "configs.html", response)
}

func make_config_list_model(model swarm.Config) configListModel {
	// utils.PrintAsJson(model)
	return configListModel{
		ID:      model.ID,
		Name:    model.Spec.Name,
		Created: model.CreatedAt.String(),
		Updated: model.UpdatedAt.String(),
	}
}
