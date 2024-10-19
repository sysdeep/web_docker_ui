package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) ActionSecretRemove(c echo.Context) error {
	name := c.Param("name")
	h.logger.Info("remove secret: " + name)

	err := h.docker_client.SecretRemove(context.Background(), name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.Redirect(http.StatusFound, "/secrets")
}
