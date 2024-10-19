package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/labstack/echo/v4"
)

// containerState stores container's running state
// it's part of ContainerJSONBase and will return by "inspect" command
type containerState struct {
	Status string `json:"status"` // String representation of the container state. Can be one of "created", "running", "paused", "restarting", "removing", "exited", or "dead"
	// Running    bool
	// Paused     bool
	// Restarting bool
	// OOMKilled  bool
	// Dead       bool
	// Pid        int
	// ExitCode   int
	// Error      string
	StartedAt string `json:"started"`
	// FinishedAt string
	// Health     *Health `json:",omitempty"`
}

// mountPoint represents a mount point configuration inside the container.
// This is used for reporting the mountpoints in use by a container.
type mountPoint struct {
	// Type is the type of mount, see `Type<foo>` definitions in
	// github.com/docker/docker/api/types/mount.Type
	// Type mount.Type `json:",omitempty"`

	// Name is the name reference to the underlying data defined by `Source`
	// e.g., the volume name.
	Name string `json:"name"`

	// Source is the source location of the mount.
	//
	// For volumes, this contains the storage location of the volume (within
	// `/var/lib/docker/volumes/`). For bind-mounts, and `npipe`, this contains
	// the source (host) part of the bind-mount. For `tmpfs` mount points, this
	// field is empty.
	// Source string

	// Destination is the path relative to the container root (`/`) where the
	// Source is mounted inside the container.
	Destination string `json:"destination"`

	// Driver is the volume driver used to create the volume (if it is a volume).
	// Driver string `json:",omitempty"`

	// Mode is a comma separated list of options supplied by the user when
	// creating the bind/volume mount.
	//
	// The default is platform-specific (`"z"` on Linux, empty on Windows).
	// Mode string

	// RW indicates whether the mount is mounted writable (read-write).
	// RW bool

	// Propagation describes how mounts are propagated from the host into the
	// mount point, and vice-versa. Refer to the Linux kernel documentation
	// for details:
	// https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt
	//
	// This field is not used on Windows.
	// Propagation mount.Propagation
}

// основные данные о контейнере
type containerModel struct {
	ID      string `json:"id"`
	Created string `json:"created"`
	// Path            string
	// Args            []string
	// State           *ContainerState
	Image string `json:"image"`
	// ResolvConfPath  string
	// HostnamePath    string
	// HostsPath       string
	// LogPath         string
	// Node            *ContainerNode `json:",omitempty"` // Node is only propagated by Docker Swarm standalone API
	Name         string `json:"name"`
	RestartCount int    `json:"restart_count"`
	// Driver          string
	// Platform        string
	// MountLabel      string
	// ProcessLabel    string
	// AppArmorProfile string
	// ExecIDs         []string
	// HostConfig      *container.HostConfig
	// GraphDriver     GraphDriverData
	// SizeRw          *int64 `json:",omitempty"`
	// SizeRootFs      *int64 `json:",omitempty"`
}

// Config contains the configuration data about a container.
// It should hold only portable information about the container.
// Here, "portable" means "independent from the host we are running on".
// Non-portable information *should* appear in HostConfig.
// All fields added to this struct must be marked `omitempty` to keep getting
// predictable hashes from the old `v1Compatibility` configuration.
type containerConfig struct {
	// Hostname        string              // Hostname
	// Domainname      string              // Domainname
	// User            string              // User that will run the command(s) inside the container, also support user:group
	// AttachStdin     bool                // Attach the standard input, makes possible user interaction
	// AttachStdout    bool                // Attach the standard output
	// AttachStderr    bool                // Attach the standard error
	// ExposedPorts    nat.PortSet         `json:",omitempty"` // List of exposed ports
	// Tty             bool                // Attach standard streams to a tty, including stdin if it is not closed.
	// OpenStdin       bool                // Open stdin
	// StdinOnce       bool                // If true, close stdin after the 1 attached client disconnects.
	Env []string `json:"env"` // List of environment variable to set in the container
	Cmd string   `json:"cmd"` // Command to run when starting the container
	// Healthcheck     *HealthConfig       `json:",omitempty"` // Healthcheck describes how to check the container is healthy
	// ArgsEscaped     bool                `json:",omitempty"` // True if command is already escaped (meaning treat as a command line) (Windows specific).
	Image string `json:"image"` // Name of the image as it was passed by the operator (e.g. could be symbolic)
	// Volumes         map[string]struct{} // List of volumes (mounts) used for the container
	// WorkingDir      string              // Current directory (PWD) in the command will be launched
	Entrypoint string `json:"entrypoint"` // Entrypoint to run when starting the container
	// NetworkDisabled bool                `json:",omitempty"` // Is network disabled

	// Mac Address of the container.
	//
	// Deprecated: this field is deprecated since API v1.44. Use EndpointSettings.MacAddress instead.
	// MacAddress  string            `json:",omitempty"`
	// OnBuild     []string          // ONBUILD metadata that were defined on the image Dockerfile
	// Labels      map[string]string // List of labels set to this container
	// StopSignal  string            `json:",omitempty"` // Signal to stop a container
	// StopTimeout *int              `json:",omitempty"` // Timeout (in seconds) to stop a container
	// Shell       strslice.StrSlice `json:",omitempty"` // Shell for shell-form of RUN, CMD, ENTRYPOINT
}

// PortBinding represents a binding between a Host IP address and a Host Port
type PortBinding struct {
	// HostIP is the host IP Address
	HostIP string `json:"host_ip"`
	// HostPort is the host port number
	HostPort string `json:"host_port"`
}

// PortMap is a collection of PortBinding indexed by Port
type PortMap map[Port][]PortBinding

// PortSet is a collection of structs indexed by Port
// type PortSet map[Port]struct{}

// Port is a string containing port number and protocol in the format "80/tcp"
type Port string

// networkEndpointSettings stores the network endpoint details
type networkEndpointSettings struct {
	// Configurations
	// IPAMConfig *EndpointIPAMConfig
	// Links      []string
	// Aliases    []string // Aliases holds the list of extra, user-specified DNS names for this endpoint.

	// MacAddress may be used to specify a MAC address when the container is created.
	// Once the container is running, it becomes operational data (it may contain a
	// generated address).
	MacAddress string `json:"mac_address"`
	// Operational data
	NetworkID string `json:"network_id"`
	// EndpointID          string
	Gateway   string `json:"gateway"`
	IPAddress string `json:"ip_address"`
	// IPPrefixLen         int
	// IPv6Gateway         string
	// GlobalIPv6Address   string
	// GlobalIPv6PrefixLen int
	// DriverOpts          map[string]string
	// DNSNames holds all the (non fully qualified) DNS names associated to this endpoint. First entry is used to
	// generate PTR records.
	// DNSNames []string
}

// networkSettings exposes the network settings in the api
type networkSettings struct {
	// NetworkSettingsBase
	Ports PortMap `json:"ports"` // Ports is a collection of PortBinding indexed by Port
	// DefaultNetworkSettings
	Networks map[string]networkEndpointSettings `json:"networks"`
}

type containerPageModel struct {
	Container containerModel `json:"container"`
	State     containerState `json:"state"`

	Mounts  []mountPoint    `json:"mounts"`
	Config  containerConfig `json:"config"`
	Network networkSettings `json:"network"`
}

func (h *Api) GetContainer(c echo.Context) error {
	container_id := c.Param("id")

	inspect_data, err := h.docker_client.ContainerInspect(context.Background(), container_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// fmt.Printf("%+v\n", inspect_data)
	// fmt.Printf("network: \n\t%+v\n", inspect_data.NetworkSettings)
	// fmt.Printf("networks: \n\t%+v\n", inspect_data.NetworkSettings.Networks)

	var mount_points []mountPoint
	for _, mp := range inspect_data.Mounts {
		mount_points = append(mount_points, make_container_mount_point_view(mp))
	}

	response := containerPageModel{

		Container: make_container_model(inspect_data.ContainerJSONBase),
		State:     make_container_state_view(inspect_data.State),

		Mounts:  mount_points,
		Config:  make_container_config_view(inspect_data.Config),
		Network: make_network_settings_view(inspect_data.NetworkSettings),
	}
	// fmt.Printf("\n%+v\n", response)

	return c.JSON(http.StatusOK, response)
}

func make_container_model(data *types.ContainerJSONBase) containerModel {
	return containerModel{
		ID:           data.ID,
		Name:         data.Name,
		Image:        data.Image,
		Created:      data.Created,
		RestartCount: data.RestartCount,
	}
}

func make_container_state_view(state *types.ContainerState) containerState {
	return containerState{
		Status: state.Status,
		// Running    bool
		// Paused     bool
		// Restarting bool
		// OOMKilled  bool
		// Dead       bool
		// Pid        int
		// ExitCode   int
		// Error      string
		StartedAt: state.StartedAt,
		// FinishedAt string
	}
}

func make_container_mount_point_view(point types.MountPoint) mountPoint {
	// fmt.Printf("\n%+v\n", point)
	return mountPoint{
		Name:        point.Name,
		Destination: point.Destination,
	}
}

func make_container_config_view(data *container.Config) containerConfig {
	// fmt.Printf("\n%+v\n", data)
	return containerConfig{
		Env:        data.Env,
		Image:      data.Image,
		Cmd:        strings.Join(data.Cmd, " "),
		Entrypoint: strings.Join(data.Entrypoint, " "),
	}
}

func make_network_settings_view(settings *types.NetworkSettings) networkSettings {

	// convert ports
	ports_map := make(PortMap)
	for key, value := range settings.Ports {
		// fmt.Println(key, value)

		var ports []PortBinding
		for _, pb := range value {
			ports = append(ports, PortBinding{HostIP: pb.HostIP, HostPort: pb.HostPort})
		}
		ports_map[Port(key)] = ports
	}

	// convert endpoints
	networks_map := make(map[string]networkEndpointSettings)
	for ep, ep_settings := range settings.Networks {
		networks_map[ep] = networkEndpointSettings{
			Gateway:    ep_settings.Gateway,
			IPAddress:  ep_settings.IPAddress,
			MacAddress: ep_settings.MacAddress,
			NetworkID:  ep_settings.NetworkID,
		}
	}

	return networkSettings{
		Ports:    ports_map,
		Networks: networks_map,
	}
}
