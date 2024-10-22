package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

// handler
func (h *Api) GetContainerStats(c echo.Context) error {

	id := c.Param("id")

	stats_result, err := h.docker_client.ContainerStatsOneShot(context.Background(), id)

	if err != nil {
		slog.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// TODO: ???????
	result := []byte{}
	stats_result.Body.Read(result)
	fmt.Println(result)

	return c.JSON(http.StatusOK, stats_result)
	// return c.JSON(http.StatusOK, newGetContainerTopResponse(top_result))
}

// type getContainerTopResponse struct {
// 	Top containerTopResult `json:"top"`
// }
//
// func newGetContainerTopResponse(model container.ContainerTopOKBody) getContainerTopResponse {
// 	top := containerTopResult{
// 		Processes: model.Processes,
// 		Titles:    model.Titles,
// 	}
//
// 	return getContainerTopResponse{
// 		Top: top,
// 	}
// }
//
// type containerTopResult struct {
// 	// Required: true
// 	Processes [][]string `json:"processes"`
//
// 	// The ps column titles
// 	// Required: true
// 	Titles []string `json:"titles"`
// }
