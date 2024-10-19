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
type secretListModel struct {
	ID      string
	Name    string
	Created string
	Updated string
}

type secretsPageModel struct {
	Secrets []secretListModel
	Total   int
}

// handler
func (h *Handlers) SecretsPage(c echo.Context) error {
	secrets_data, err := h.docker_client.SecretList(context.Background(), types.SecretListOptions{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var secrets []secretListModel
	for _, v := range secrets_data {
		secrets = append(secrets, make_secret_list_model(v))
	}

	sort.SliceStable(secrets, func(i, j int) bool {
		return secrets[i].Name < secrets[j].Name
	})

	response := secretsPageModel{
		Secrets: secrets,
		Total:   len(secrets),
	}

	return c.Render(http.StatusOK, "secrets.html", response)
}

func make_secret_list_model(model swarm.Secret) secretListModel {
	// utils.PrintAsJson(model)
	return secretListModel{
		ID:      model.ID,
		Name:    model.Spec.Name,
		Created: model.CreatedAt.String(),
		Updated: model.UpdatedAt.String(),
	}
}
