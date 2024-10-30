package api

import (
	"context"
	"hdu/internal/webserver/api/docker_adapter"
	"hdu/internal/webserver/api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// handler --------------------------------------------------------------------
func (h *Api) GetSecret(c echo.Context) error {
	id := c.Param("id")
	secret_data, _, err := h.docker_client.SecretInspectWithRaw(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	// utils.PrintAsJson(secret_data)

	response := secretPageModel{
		Secret: docker_adapter.NewSecretModel(secret_data),
	}

	return c.JSON(http.StatusOK, response)
}

// models ---------------------------------------------------------------------
type secretPageModel struct {
	Secret models.Secret `json:"secret"`
	// Total   int
}
