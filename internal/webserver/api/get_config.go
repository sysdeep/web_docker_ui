package api

import (
	"context"
	"hdu/internal/utils"
	"net/http"

	"github.com/docker/docker/api/types/swarm"
	"github.com/labstack/echo/v4"
)

// models
type configModel struct {
	ID       string `json:"id"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
	Name     string `json:"name"`
	DataText string `json:"data_text"`
}

type configPageModel struct {
	Config configModel `json:"config"`
}

// handler
func (h *Api) GetConfig(c echo.Context) error {
	id := c.Param("id")
	config_data, _, err := h.docker_client.ConfigInspectWithRaw(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	response := configPageModel{
		Config: make_config_model(config_data),
	}

	return c.JSON(http.StatusOK, response)
}

func make_config_model(model swarm.Config) configModel {
	// utils.PrintAsJson(model)
	config_data := string(model.Spec.Data)
	if utils.IsJSON(config_data) {
		config_data = utils.FormatJson(config_data, "    ")
	}

	return configModel{
		ID:       model.ID,
		Name:     model.Spec.Name,
		Created:  model.CreatedAt.String(),
		Updated:  model.UpdatedAt.String(),
		DataText: config_data,
	}
}
