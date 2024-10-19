package services

import (
	"context"
	"sort"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

type VolumesService struct {
	docker_client *client.Client
}

func NewVolumesService(docker_client *client.Client) *VolumesService {
	return &VolumesService{docker_client: docker_client}
}

func (cs *VolumesService) GetAll() ([]VolumeListModel, error) {
	var result []VolumeListModel

	// used volumes
	used_filter := filters.KeyValuePair{Key: "dangling", Value: "false"}
	used_volumes_list, err := cs.docker_client.VolumeList(context.Background(), volume.ListOptions{Filters: filters.NewArgs(used_filter)})
	if err != nil {
		return result, err
	}

	for _, c := range used_volumes_list.Volumes {
		result = append(result, make_volume_list_model(c, true))
	}

	// non used volumes
	non_used_filter := filters.KeyValuePair{Key: "dangling", Value: "true"}
	non_used_volumes_list, err := cs.docker_client.VolumeList(context.Background(), volume.ListOptions{Filters: filters.NewArgs(non_used_filter)})
	if err != nil {
		return result, err
	}

	for _, c := range non_used_volumes_list.Volumes {
		result = append(result, make_volume_list_model(c, false))
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result, nil
}

// models ---------------------------------------------------------------------

type VolumeListModel struct {
	Name string

	// cluster volume
	// ClusterVolume *ClusterVolume `json:"ClusterVolume,omitempty"`

	// Date/Time the volume was created.
	CreatedAt string

	// Name of the volume driver used by the volume.
	Driver string

	// User-defined key/value metadata.
	// Required: true
	// Labels map[string]string `json:"Labels"`

	// Mount path of the volume on the host.
	Mountpoint string

	// The driver specific options used when creating the volume.
	//
	// Required: true
	// Options map[string]string `json:"Options"`

	// The level at which the volume exists. Either `global` for cluster-wide,
	// or `local` for machine level.
	//
	// Required: true
	// Scope string `json:"Scope"`

	// Low-level details about the volume, provided by the volume driver.
	// Details are returned as a map with key/value pairs:
	// `{"key":"value","key2":"value2"}`.
	//
	// The `Status` field is optional, and is omitted if the volume driver
	// does not support this feature.
	//
	// Status map[string]interface{} `json:"Status,omitempty"`

	// usage data
	// UsageData *UsageData `json:"UsageData,omitempty"`

	StackName string
	Used      bool
}

const volume_stack_label = "com.docker.stack.namespace"

func make_volume_list_model(data *volume.Volume, used bool) VolumeListModel {

	stack_name := ""
	if stack_name_labeled, ok := data.Labels[volume_stack_label]; ok {
		stack_name = stack_name_labeled
	}

	return VolumeListModel{
		Name:       data.Name,
		CreatedAt:  data.CreatedAt,
		Driver:     data.Driver,
		Mountpoint: data.Mountpoint,
		StackName:  stack_name,
		Used:       used,
	}

}
