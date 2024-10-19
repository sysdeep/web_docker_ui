package api

import (
	"context"
	"net/http"
	"time"

	"github.com/docker/docker/api/types/image"
	"github.com/labstack/echo/v4"
)

// models
type imageListModel struct {

	// Number of containers using this image. Includes both stopped and running
	// containers.
	//
	// This size is not calculated by default, and depends on which API endpoint
	// is used. `-1` indicates that the value has not been set / calculated.
	//
	// Required: true
	Containers int64 `json:"containers"`

	Created string `json:"created"`

	// ID is the content-addressable ID of an image.
	//
	// This identifier is a content-addressable digest calculated from the
	// image's configuration (which includes the digests of layers used by
	// the image).
	//
	// Note that this digest differs from the `RepoDigests` below, which
	// holds digests of image manifests that reference the image.
	//
	// Required: true
	ID string `json:"id"`

	// User-defined key/value metadata.
	// Required: true
	// Labels map[string]string `json:"Labels"`

	// ID of the parent image.
	//
	// Depending on how the image was created, this field may be empty and
	// is only set for images that were built/created locally. This field
	// is empty if the image was pulled from an image registry.
	//
	// Required: true
	// ParentID string `json:"ParentId"`

	// List of content-addressable digests of locally available image manifests
	// that the image is referenced from. Multiple manifests can refer to the
	// same image.
	//
	// These digests are usually only available if the image was either pulled
	// from a registry, or if the image was pushed to a registry, which is when
	// the manifest is generated and its digest calculated.
	//
	// Required: true
	// RepoDigests []string `json:"RepoDigests"`

	// List of image names/tags in the local image cache that reference this
	// image.
	//
	// Multiple image tags can refer to the same image, and this list may be
	// empty if no tags reference the image, in which case the image is
	// "untagged", in which case it can still be referenced by its ID.
	RepoTags []string `json:"tags"`

	// Total size of image layers that are shared between this image and other
	// images.
	//
	// This size is not calculated by default. `-1` indicates that the value
	// has not been set / calculated.
	//
	// Required: true
	// SharedSize int64 `json:"SharedSize"`

	// Total size of the image including all layers it is composed of.
	Size int64 `json:"size"`

	// Total size of the image including all layers it is composed of.
	//
	// Deprecated: this field is omitted in API v1.44, but kept for backward compatibility. Use Size instead.
	// VirtualSize int64 `json:"VirtualSize,omitempty"`

}

type imagesPageModel struct {
	Images []imageListModel `json:"images"`
	Total  int              `json:"total"`
}

// handler
func (h *Api) GetImages(c echo.Context) error {
	images_list, err := h.docker_client.ImageList(context.Background(), image.ListOptions{All: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var images []imageListModel
	for _, image := range images_list {
		images = append(images, make_image_list_model(image))
	}

	response := imagesPageModel{
		Images: images,
		Total:  len(images),
	}

	return c.JSON(http.StatusOK, response)
}

// make model
func make_image_list_model(data image.Summary) imageListModel {
	// fmt.Printf("\n%+v\n", data)

	created_time := time.Unix(data.Created, 0)
	created_string := created_time.Format("2006-01-02 15:04:05")

	return imageListModel{
		Containers: data.Containers,
		Created:    created_string,
		ID:         data.ID,
		RepoTags:   data.RepoTags,
		Size:       data.Size,
	}
}
