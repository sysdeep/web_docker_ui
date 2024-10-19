package api

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/api/types/system"
	"github.com/labstack/echo/v4"
)

// main page model
type mainPageModel struct {
	DaemonHost    string     `json:"daemon_host"`    // DaemonHost returns the host address used by the client
	ClientVersion string     `json:"client_version"` // the API version used by this client
	SystemInfo    systemInfo `json:"system"`
}

func (h *Api) GetInfo(c echo.Context) error {

	// слишком долго
	// disk_usage, err := h.docker_client.DiskUsage(context.Background(), types.DiskUsageOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", disk_usage)

	sys_info, err := h.docker_client.Info(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	response := mainPageModel{
		DaemonHost:    h.docker_client.DaemonHost(),
		ClientVersion: h.docker_client.ClientVersion(),
		SystemInfo:    make_system_info(sys_info),
	}

	return c.JSON(http.StatusOK, response)
}

// systemInfo
type systemInfo struct {
	ID                string `json:"id"`
	Containers        int    `json:"containers"`
	ContainersRunning int    `json:"containers_running"`
	ContainersPaused  int    `json:"containers_paused"`
	ContainersStopped int    `json:"containers_stopped"`
	Images            int    `json:"images"`
	Driver            string `json:"driver"`
	// DriverStatus       [][2]string
	// SystemStatus       [][2]string `json:",omitempty"` // SystemStatus is only propagated by the Swarm standalone API
	// Plugins            PluginsInfo
	// MemoryLimit        bool
	// SwapLimit          bool
	// KernelMemory       bool `json:",omitempty"` // Deprecated: kernel 5.4 deprecated kmem.limit_in_bytes
	// KernelMemoryTCP    bool `json:",omitempty"` // KernelMemoryTCP is not supported on cgroups v2.
	// CPUCfsPeriod       bool `json:"CpuCfsPeriod"`
	// CPUCfsQuota        bool `json:"CpuCfsQuota"`
	// CPUShares          bool
	// CPUSet             bool
	// PidsLimit          bool
	// IPv4Forwarding     bool
	// BridgeNfIptables   bool
	// BridgeNfIP6tables  bool `json:"BridgeNfIp6tables"`
	// Debug              bool
	// NFd                int
	// OomKillDisable     bool
	// NGoroutines        int
	// SystemTime         string
	// LoggingDriver      string
	// CgroupDriver       string
	// CgroupVersion      string `json:",omitempty"`
	// NEventsListener    int
	KernelVersion   string `json:"kernel_version"`
	OperatingSystem string `json:"operating_system"`
	OSVersion       string `json:"os_version"`
	OSType          string `json:"os_type"`
	// Architecture       string
	// IndexServerAddress string
	// RegistryConfig     *registry.ServiceConfig
	NCPU     int   `json:"ncpu"`
	MemTotal int64 `json:"mem_total"`
	// GenericResources   []swarm.GenericResource
	// DockerRootDir      string
	// HTTPProxy          string `json:"HttpProxy"`
	// HTTPSProxy         string `json:"HttpsProxy"`
	// NoProxy            string
	Name string `json:"name"`
	// Labels             []string
	// ExperimentalBuild  bool
	ServerVersion string `json:"server_version"` // docker engine server version
	// Runtimes           map[string]RuntimeWithStatus
	DefaultRuntime string    `json:"default_runtime"` // runc
	Swarm          swarmInfo `json:"swarm"`
	// // LiveRestoreEnabled determines whether containers should be kept
	// // running when the daemon is shutdown or upon daemon start if
	// // running containers are detected
	// LiveRestoreEnabled  bool
	// Isolation           container.Isolation
	// InitBinary          string
	// ContainerdCommit    Commit
	// RuncCommit          Commit
	// InitCommit          Commit
	// SecurityOptions     []string
	// ProductLicense      string               `json:",omitempty"`
	DefaultAddressPools []networkAddressPool `json:"default_addresses_pools"`
	// CDISpecDirs         []string
	//
	// // Legacy API fields for older API versions.
	// legacyFields
	//
	// // Warnings contains a slice of warnings that occurred  while collecting
	// // system information. These warnings are intended to be informational
	// // messages for the user, and are not intended to be parsed / used for
	// // other purposes, as they do not have a fixed format.
	// Warnings []string
}

// NetworkAddressPool is a temp struct used by [Info] struct.
type networkAddressPool struct {
	Base string `json:"base"`
	Size int    `json:"size"`
}

// Info represents generic information about swarm.
type swarmInfo struct {
	NodeID   string `json:"node_id"`
	NodeAddr string `json:"node_addr"`

	// LocalNodeState   LocalNodeState
	// ControlAvailable bool
	// Error            string
	//
	// RemoteManagers []Peer
	// Nodes          int `json:",omitempty"`
	// Managers       int `json:",omitempty"`
	//
	// Cluster *ClusterInfo `json:",omitempty"`
	//
	// Warnings []string `json:",omitempty"`
}

func make_swarm_info(data swarm.Info) swarmInfo {
	return swarmInfo{
		NodeID:   data.NodeID,
		NodeAddr: data.NodeAddr,
	}
}

func make_system_info(data system.Info) systemInfo {
	var network_pool []networkAddressPool
	for _, np := range data.DefaultAddressPools {
		network_pool = append(network_pool, networkAddressPool{Base: np.Base, Size: np.Size})
	}

	return systemInfo{
		ID:                  data.ID,
		Containers:          data.Containers,
		ContainersRunning:   data.ContainersRunning,
		ContainersPaused:    data.ContainersPaused,
		ContainersStopped:   data.ContainersStopped,
		Images:              data.Images,
		Driver:              data.Driver,
		KernelVersion:       data.KernelVersion,
		OperatingSystem:     data.OperatingSystem,
		OSVersion:           data.OSVersion,
		OSType:              data.OSType,
		NCPU:                data.NCPU,
		MemTotal:            data.MemTotal,
		DefaultAddressPools: network_pool,
		DefaultRuntime:      data.DefaultRuntime,
		ServerVersion:       data.ServerVersion,
		Name:                data.Name,
		Swarm:               make_swarm_info(data.Swarm),
	}
}
