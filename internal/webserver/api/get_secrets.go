package api

import (
	"context"
	"hdu/internal/webserver/api/docker_adapter"
	"hdu/internal/webserver/api/models"
	"net/http"
	"sort"

	"github.com/docker/docker/api/types"
	"github.com/labstack/echo/v4"
)

// handler --------------------------------------------------------------------
func (h *Api) GetSecrets(c echo.Context) error {
	secrets_data, err := h.docker_client.SecretList(context.Background(), types.SecretListOptions{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	secrets := []models.Secret{}
	for _, v := range secrets_data {
		secrets = append(secrets, docker_adapter.NewSecretModel(v))
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

// models ---------------------------------------------------------------------
type secretsPageModel struct {
	Secrets []models.Secret `json:"secrets"`
	Total   int             `json:"total"`
}
