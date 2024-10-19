package api

import (
	"hdu/internal/registry_client"
	"net/http"

	"github.com/labstack/echo/v4"
)

// handler
func (h *Api) GetRegistryRepository(c echo.Context) error {

	id := c.Param("id")
	repo, err := h.registry_client.GetRepository(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newRegistryRepositoryResponse(repo))
}

type registryRepositoryResponse struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func newRegistryRepositoryResponse(model registry_client.RepositoryModel) registryRepositoryResponse {
	return registryRepositoryResponse{
		ID:   model.ID,
		Name: model.Name,
		Tags: model.Tags,
	}
}
