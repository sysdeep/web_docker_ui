package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// handler
func (h *Api) RemoveRegistryRepositoryTag(c echo.Context) error {

	id := c.Param("id")
	tag := c.Param("tag")

	manifest, err := h.registry_client.GetManivestV2(id, tag)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	err = h.registry_client.RemoveManifest(id, manifest.ContentDigest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, stdResponse{status: true})
}
