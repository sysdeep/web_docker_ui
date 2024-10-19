package api

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
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

type secretsPageModel struct {
	Secrets []secretListModel `json:"secrets"`
	Total   int               `json:"total"`
}

// handler
func (h *Api) GetSecrets(c echo.Context) error {
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

	return c.JSON(http.StatusOK, response)
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
