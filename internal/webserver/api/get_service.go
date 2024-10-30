package api

import (
	"context"
	"hdu/internal/utils"
	"hdu/internal/webserver/api/domain/service_model"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/labstack/echo/v4"
)

// TODO: модель данных сервиса та же самая что и для списка
// handler --------------------------------------------------------------------
func (h *Api) GetService(c echo.Context) error {
	id := c.Param("id")

	options := types.ServiceInspectOptions{}
	service_model_d, raw, err := h.docker_client.ServiceInspectWithRaw(context.Background(), id, options)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	utils.PrintAsJson(service_model_d)
	utils.PrintAsJson(raw)

	response := service_model.NewServiceListModel(service_model_d)

	return c.JSON(http.StatusOK, response)
}

// models ---------------------------------------------------------------------
