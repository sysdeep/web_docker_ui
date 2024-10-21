package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/labstack/echo/v4"
)

// handler
func (h *Api) GetContainerTop(c echo.Context) error {

	id := c.Param("id")

	top_options := []string{}
	top_result, err := h.docker_client.ContainerTop(context.Background(), id, top_options)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newGetContainerTopResponse(top_result))
}

type getContainerTopResponse struct {
	Top containerTopResult `json:"top"`
}

func newGetContainerTopResponse(model container.ContainerTopOKBody) getContainerTopResponse {
	top := containerTopResult{
		Processes: model.Processes,
		Titles:    model.Titles,
	}

	return getContainerTopResponse{
		Top: top,
	}
}

type containerTopResult struct {
	// Required: true
	Processes [][]string `json:"processes"`

	// The ps column titles
	// Required: true
	Titles []string `json:"titles"`
}
