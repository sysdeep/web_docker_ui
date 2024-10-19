package containers_page

import "github.com/visualfc/atk/tk"

type actionsBar struct {
	*tk.Frame
	vm *ContainersVM
	// images_service *services.ImagesService
	// list_frame     *ImagesFrame
}

func newActionsBar(parent tk.Widget, vm *ContainersVM) *actionsBar {
	fr := tk.NewFrame(parent)

	btn_refresh := tk.NewButton(fr, "refresh")

	layout := tk.NewHPackLayout(fr)

	layout.AddWidget(tk.NewLayoutSpacer(fr, 1, true))
	layout.AddWidget(btn_refresh,
		tk.PackAttrSideRight(),
	)

	// events -------------------------------------------------------------------
	btn_refresh.OnCommand(func() {
		vm.refresh()
	})

	return &actionsBar{
		Frame: fr,
		vm:    vm,
	}

}
