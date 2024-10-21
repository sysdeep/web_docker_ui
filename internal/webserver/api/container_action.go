package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
)

// handler
func (h *Api) ContainerAction(c echo.Context) error {

	var request containerActionRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var action_error error
	switch request.Action {
	case "stop":
		action_error = actionStopContainer(h.docker_client, request.ID)
	case "start":
		action_error = actionStartContainer(h.docker_client, request.ID)
	}

	if action_error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, action_error)
	}
	// action := c.Param("action")

	// // Получение списка запуцщенных контейнеров(docker ps)
	// raw_containers, err := h.docker_client.ContainerList(context.Background(), container.ListOptions{All: true})
	//
	// utils.PrintAsJson(raw_containers)
	// containers := []containerListModel{}
	// for _, c := range raw_containers {
	// 	containers = append(containers, convert_container(c))
	// }
	//
	// response := containersPageModel{
	// 	Containers: containers,
	// }
	//
	return c.JSON(http.StatusOK, request)
}

type containerActionRequest struct {
	Action string `json:"action"`
	ID     string `json:"id"`
}

func actionStopContainer(docker_client *client.Client, id string) error {
	slog.Info("container action stop: " + id)

	// NOTE: example top
	// top_options := []string{}
	// aaa, err := docker_client.ContainerTop(context.Background(), id, top_options)
	// fmt.Println(aaa)
	// fmt.Println(err)

	stop_options := container.StopOptions{}
	return docker_client.ContainerStop(context.Background(), id, stop_options)
}

func actionStartContainer(docker_client *client.Client, id string) error {
	slog.Info("container action start: " + id)

	start_options := container.StartOptions{}
	return docker_client.ContainerStart(context.Background(), id, start_options)
}
