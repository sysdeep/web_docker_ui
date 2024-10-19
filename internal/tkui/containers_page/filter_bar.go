package containers_page

import (
	"fmt"

	"github.com/visualfc/atk/tk"
)

type filterBar struct {
	*tk.Frame
	vm *ContainersVM
}

func newFilterBar(parent tk.Widget, vm *ContainersVM) *filterBar {
	fr := tk.NewFrame(parent)
	bar := &filterBar{
		Frame: fr,
		vm:    vm,
	}

	// status
	st_group := tk.NewRadioGroup()
	st_all := tk.NewRadioButton(fr, "all")
	st_running := tk.NewRadioButton(fr, "running")
	st_stopped := tk.NewRadioButton(fr, "stopped")
	st_group.AddRadios(st_all, st_running, st_stopped)

	current_filter := vm.get_filter()
	// NOTE: несколько SetChecked нельзя, ломается вся логика... поэтому switch
	switch current_filter.status {
	case FILTER_ST_ALL:
		st_all.SetChecked(true)
	case FILTER_ST_RUNNING:
		st_running.SetChecked(true)
	case FILTER_ST_STOPPED:
		st_stopped.SetChecked(true)
	}

	// layout
	main_layout := tk.NewHPackLayout(fr)
	main_layout.AddWidgets(st_all)
	main_layout.AddWidgets(st_running)
	main_layout.AddWidgets(st_stopped)

	st_group.OnRadioChanged(func() {
		w := st_group.CheckedRadio()
		fmt.Println(w)

		if w == st_all {
			vm.set_filter_status(FILTER_ST_ALL)
		}
		if w == st_running {
			vm.set_filter_status(FILTER_ST_RUNNING)
		}
		if w == st_stopped {
			vm.set_filter_status(FILTER_ST_STOPPED)
		}

	})

	return bar
}
