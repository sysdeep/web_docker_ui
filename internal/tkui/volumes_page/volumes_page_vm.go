package volumes_page

import "hdu/internal/services"

const (
	EV_REFRESH = "refresh"
)

type VolumesPageVM struct {
	observers       []Observer
	volumes_service *services.VolumesService
}

func NewVolumesPageVM(volumes_service *services.VolumesService) *VolumesPageVM {

	return &VolumesPageVM{
		volumes_service: volumes_service,
	}
}

func (vm *VolumesPageVM) get_volumes() []services.VolumeListModel {
	volumes, _ := vm.volumes_service.GetAll()
	return volumes
}

func (vm *VolumesPageVM) register(observer Observer) {
	vm.observers = append(vm.observers, observer)
}

func (vm *VolumesPageVM) refresh() {
	vm.send_event(EV_REFRESH)
}

func (vm *VolumesPageVM) send_event(ev string) {
	for _, observer := range vm.observers {
		observer.update(ev)
	}
}

type Observer interface {
	update(ev string)
}
