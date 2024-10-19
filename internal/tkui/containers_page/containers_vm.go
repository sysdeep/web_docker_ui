package containers_page

import (
	"fmt"
	"hdu/internal/services"
)

const (
	EV_REFRESH        = "ev_refresh"
	EV_FILTER         = "ev_filter"
	FILTER_ST_ALL     = "filter_st_all"
	FILTER_ST_RUNNING = "filter_st_running"
	FILTER_ST_STOPPED = "filter_st_stopped"
)

type containersFilter struct {
	status string
}

type ContainersVM struct {
	observers            []Observer
	containers_service   *services.ContainersService
	root_actions_handler RootActionsHandler
	filter               containersFilter
}

func NewContainersVM(containers_service *services.ContainersService, root_actions_handler RootActionsHandler) *ContainersVM {

	filter := containersFilter{
		status: FILTER_ST_ALL,
	}

	return &ContainersVM{
		containers_service:   containers_service,
		root_actions_handler: root_actions_handler,
		filter:               filter,
	}
}

func (vm *ContainersVM) get_containers() []services.ContainerListModel {
	containers, _ := vm.containers_service.GetAll()

	return containers
}

func (vm *ContainersVM) refresh() {
	vm.send_event(EV_REFRESH)
}

func (vm *ContainersVM) show_container(container services.ContainerListModel) {
	fmt.Println("show container", container.Name)
	vm.root_actions_handler.ShowContainer(container.ID)
}

func (vm *ContainersVM) set_filter_status(status string) {
	vm.filter.status = status
	vm.send_event(EV_FILTER)
}

func (vm *ContainersVM) get_filter() containersFilter {
	return vm.filter
}

// observer interface ---------------------------------------------------------
func (vm *ContainersVM) register(observer Observer) {
	vm.observers = append(vm.observers, observer)
}

func (vm *ContainersVM) send_event(ev string) {
	for _, observer := range vm.observers {
		observer.update(ev)
	}
}

type Observer interface {
	update(ev string)
}

type RootActionsHandler interface {
	ShowContainer(container_id string)
}
