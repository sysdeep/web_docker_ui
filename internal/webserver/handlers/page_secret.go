package handlers

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/types/swarm"
	"github.com/labstack/echo/v4"
)

// models
type secretModel struct {
	ID      string
	Name    string
	Created string
	Updated string
}

type secretPageModel struct {
	Secret secretModel
	// Total   int
}

// handler
func (h *Handlers) SecretPage(c echo.Context) error {
	id := c.Param("id")
	secret_data, _, err := h.docker_client.SecretInspectWithRaw(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	// utils.PrintAsJson(secret_data)

	response := secretPageModel{
		Secret: make_secret_model(secret_data),
	}

	return c.Render(http.StatusOK, "secret.html", response)
}

func make_secret_model(model swarm.Secret) secretModel {
	// utils.PrintAsJson(model)
	return secretModel{
		ID:      model.ID,
		Name:    model.Spec.Name,
		Created: model.CreatedAt.String(),
		Updated: model.UpdatedAt.String(),
	}
}
