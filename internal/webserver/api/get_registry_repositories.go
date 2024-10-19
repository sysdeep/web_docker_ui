package api

import (
	"hdu/internal/registry_client"
	"net/http"

	"github.com/labstack/echo/v4"
)

// handler
func (h *Api) GetRegistryRepositories(c echo.Context) error {

	catalog, err := h.registry_client.GetCatalog(10)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newRegistryRepositoriesResponse(&catalog))
}

type registryRepositoriesResponse struct {
	Repositories []repositoryListModel `json:"repositories"`
}

func newRegistryRepositoriesResponse(model *registry_client.Catalog) registryRepositoriesResponse {

	repos := []repositoryListModel{}
	for _, row := range model.Repositories {
		repos = append(repos, newRepositoryListModel(row))
	}

	return registryRepositoriesResponse{
		Repositories: repos,
	}
}

type repositoryListModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func newRepositoryListModel(model registry_client.RepositoryListModel) repositoryListModel {
	return repositoryListModel{
		ID:   model.ID,
		Name: model.Name,
	}
}
