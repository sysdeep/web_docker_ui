package volumes_page

import (
	"github.com/visualfc/atk/tk"
)

type VolumesPage struct {
	*tk.Frame
	list_bar    *listBar
	actions_bar *actionsBar
	vm          *VolumesPageVM
}

func NewVolumesPage(parent tk.Widget, vm *VolumesPageVM) *VolumesPage {

	fr := tk.NewFrame(parent)

	page := &VolumesPage{
		Frame:       fr,
		list_bar:    newListBar(fr, vm),
		actions_bar: newActionsBar(fr, vm),
		vm:          vm,
	}

	// label
	lbl := tk.NewLabel(fr, "Volumes")

	// list

	// refresh button
	// refresh_button := tk.NewButton(fr, "Refresh")

	// lauout
	main_layout := tk.NewVPackLayout(fr)
	main_layout.AddWidget(lbl)
	main_layout.AddWidget(page.list_bar,
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
	)
	main_layout.AddWidget(page.actions_bar,
		tk.PackAttrFillX(),
	)

	// fr.BindEvent("<Enter>", func(e *tk.Event) {
	// 	fmt.Println("Enter")
	// })
	//
	// fr.BindEvent("<Activate>", func(e *tk.Event) {
	// 	fmt.Println("Activate")
	// })

	// bind events --------------------------------------------------------------

	// событие отображения(при переключении вкладок)
	fr.BindEvent("<Visibility>", func(e *tk.Event) {
		page.vm.refresh()
	})
	//
	// refresh_button.OnCommand(func() {
	// 	page.refresh()
	// })

	return page
}

// func (cp *VolumesPage) makeLayout() {
// 	// lauout
// 	main_layout := tk.NewVPackLayout(cp.Frame)
// 	// main_layout.AddWidget(lbl)
// 	main_layout.AddWidget(cp.tree,
// 		tk.PackAttrFillBoth(),
// 		tk.PackAttrExpand(true),
// 	)
// 	// main_layout.AddWidget(refresh_button)
//
// }
