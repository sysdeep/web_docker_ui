package api

import (
	"context"
	"hdu/internal/utils"
	"net/http"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/labstack/echo/v4"
)

// models

type imageContainerModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type imageModel struct {
	// holds digests of image manifests that reference the image.
	ID string `json:"id"`

	// RepoTags is a list of image names/tags in the local image cache that
	// reference this image.
	//
	// Multiple image tags can refer to the same image, and this list may be
	// empty if no tags reference the image, in which case the image is
	// "untagged", in which case it can still be referenced by its ID.
	RepoTags []string `json:"repo_tags"`

	// RepoDigests is a list of content-addressable digests of locally available
	// image manifests that the image is referenced from. Multiple manifests can
	// refer to the same image.
	//
	// These digests are usually only available if the image was either pulled
	// from a registry, or if the image was pushed to a registry, which is when
	// the manifest is generated and its digest calculated.
	// RepoDigests []string

	// Parent is the ID of the parent image.
	//
	// Depending on how the image was created, this field may be empty and
	// is only set for images that were built/created locally. This field
	// is empty if the image was pulled from an image registry.
	Parent string `json:"parent"`

	// Comment is an optional message that can be set when committing or
	// importing the image.
	Comment string `json:"comment"`

	// Created is the date and time at which the image was created, formatted in
	// RFC 3339 nano-seconds (time.RFC3339Nano).
	//
	// This information is only available if present in the image,
	// and omitted otherwise.
	Created string `json:"created"`

	// DockerVersion is the version of Docker that was used to build the image.
	//
	// Depending on how the image was created, this field may be empty.
	// DockerVersion string

	// Author is the name of the author that was specified when committing the
	// image, or as specified through MAINTAINER (deprecated) in the Dockerfile.
	// Author string
	// Config *container.Config

	// Architecture is the hardware CPU architecture that the image runs on.
	// Architecture string

	// Variant is the CPU architecture variant (presently ARM-only).
	// Variant string `json:",omitempty"`

	// OS is the Operating System the image is built to run on.
	// Os string

	// OsVersion is the version of the Operating System the image is built to
	// run on (especially for Windows).
	// OsVersion string `json:",omitempty"`

	// Size is the total size of the image including all layers it is composed of.
	Size int64 `json:"size"`

	// VirtualSize is the total size of the image including all layers it is
	// composed of.
	//
	// Deprecated: this field is omitted in API v1.44, but kept for backward compatibility. Use Size instead.
	// VirtualSize int64 `json:"VirtualSize,omitempty"`

	// GraphDriver holds information about the storage driver used to store the
	// container's and image's filesystem.
	// GraphDriver GraphDriverData

	// RootFS contains information about the image's RootFS, including the
	// layer IDs.
	// RootFS RootFS

	// Metadata of the image in the local cache.
	//
	// This information is local to the daemon, and not part of the image itself.
	// Metadata image.Metadata

}

type imageHistoryModel struct {

	// comment
	// Required: true
	// Comment string `json:"Comment"`

	Created string `json:"created"`

	// created by
	// Required: true
	// CreatedBy string `json:"CreatedBy"`

	// Id
	// Required: true
	ID string `json:"id"`

	// size
	// Required: true
	Size int64 `json:"size"`

	// tags
	// Required: true
	Tags []string `json:"tags"`
}

type imagePageModel struct {
	Image      imageModel            `json:"image"`
	History    []imageHistoryModel   `json:"history"`
	Containers []imageContainerModel `json:"containers"`
}

// handler
func (h *Api) GetImage(c echo.Context) error {
	id := c.Param("id")

	image_inspect, _, err := h.docker_client.ImageInspectWithRaw(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	history_data, err := h.docker_client.ImageHistory(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	raw_containers, err := h.docker_client.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var history []imageHistoryModel
	for _, h := range history_data {
		history = append(history, make_image_history_model(h))
	}

	image_model := make_image_model(image_inspect)
	response := imagePageModel{
		Image:      image_model,
		History:    history,
		Containers: make_image_containers(image_model, raw_containers),
	}

	// fmt.Printf("%+v\n", response)

	return c.JSON(http.StatusOK, response)
}

func make_image_history_model(data image.HistoryResponseItem) imageHistoryModel {

	created_time := time.Unix(data.Created, 0)
	created_string := created_time.Format(utils.TIME_LAYOUT)

	tags := []string{}
	if data.Tags != nil {
		tags = data.Tags
	}
	return imageHistoryModel{
		ID:      data.ID,
		Created: created_string,
		Size:    data.Size,
		Tags:    tags,
	}
}

func make_image_model(data types.ImageInspect) imageModel {

	created_time, _ := time.Parse(time.RFC3339Nano, data.Created)
	created_string := created_time.Format(utils.TIME_LAYOUT)

	return imageModel{
		ID:       data.ID,
		Size:     data.Size,
		RepoTags: data.RepoTags,
		Created:  created_string,
		Parent:   data.Parent,
		Comment:  data.Comment,
	}
}

func make_image_containers(image imageModel, containers_list []types.Container) []imageContainerModel {
	result := []imageContainerModel{}

	for _, c := range containers_list {
		if c.ImageID == image.ID {
			// NOTE: берём только 1 имя, зачем это списком сделано - непонятно...
			result = append(result, imageContainerModel{c.ID, c.Names[0]})
		}
	}
	return result
}
