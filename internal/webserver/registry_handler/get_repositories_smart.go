package registry_handler

import (
	"hdu/internal/registry_client"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

// handler
func (h *RegistryHandler) GetRegistryRepositoriesSmart(c echo.Context) error {

	catalog, err := h.registry_client.GetCatalog(10000)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	result := []registry_client.RepositoryModel{}
	for _, cat := range catalog.Repositories {
		repo, err := h.registry_client.GetRepository(cat.ID)
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		// skip empty
		// if len(repo.Tags) == 0 {
		// 	continue
		// }

		result = append(result, repo)

	}

	return c.JSON(http.StatusOK, newRegistryRepositoriesSmartResponse(result))
}

type registryRepositoriesSmartResponse struct {
	Repositories []registryRepositoryResponse `json:"repositories"`
}

func newRegistryRepositoriesSmartResponse(model []registry_client.RepositoryModel) registryRepositoriesSmartResponse {

	repos := []registryRepositoryResponse{}
	for _, row := range model {
		repos = append(repos, newRegistryRepositoryResponse(row))
	}

	return registryRepositoriesSmartResponse{
		Repositories: repos,
	}
}
