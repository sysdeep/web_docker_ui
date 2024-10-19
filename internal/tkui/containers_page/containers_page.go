package containers_page

import (
	"fmt"
	"hdu/internal/services"

	"github.com/visualfc/atk/tk"
)

type ContainersPage struct {
	*tk.Frame
	list_frame  *listBar
	actions_bar *actionsBar
	vm          *ContainersVM
}

func NewContainersPage(parent tk.Widget, vm *ContainersVM) *ContainersPage {

	fr := tk.NewFrame(parent)

	// label
	lbl := tk.NewLabel(fr, "Containers")

	filter := newFilterBar(fr, vm)

	// list
	list := newListBar(fr, vm)

	// fr.BindEvent("<Enter>", func(e *tk.Event) {
	// 	fmt.Println("Enter")
	// })
	//
	// fr.BindEvent("<Activate>", func(e *tk.Event) {
	// 	fmt.Println("Activate")
	// })

	page := &ContainersPage{
		Frame:       fr,
		list_frame:  list,
		actions_bar: newActionsBar(fr, vm),
		vm:          vm,
	}

	// layout
	main_layout := tk.NewVPackLayout(fr)
	main_layout.AddWidget(lbl)
	main_layout.AddWidget(filter,
		tk.PackAttrFillX(),
	)
	main_layout.AddWidget(list,
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
	)
	main_layout.AddWidget(page.actions_bar,
		tk.PackAttrFillX(),
	)

	// bind events --------------------------------------------------------------

	// событие отображения(при переключении вкладок)
	fr.BindEvent("<Visibility>", func(e *tk.Event) {
		vm.refresh()
	})

	// list.ConnectContainerSelected(page.OnContainerSelected)

	return page
}

// func (cp *ContainersPage) refresh() {
// 	items := cp.vm.get_containers()
//
// 	cp.list_frame.SetItems(items)
// }

func (cp *ContainersPage) OnContainerSelected(model *services.ContainerListModel) {
	fmt.Println(model)

	// TODO: in new type
	top := tk.NewWindow()

	// view := NewContainerView(top, NewFakeContainerProvider())
	// layout := tk.NewVPackLayout(top)
	// layout.AddWidget(view, tk.PackAttrFillBoth(), tk.PackAttrExpand(true))

	top.SetTitle("Container view")
	top.ShowNormal()
}

// TODO: only for tests
type FakeContainerProvider struct{}

func NewFakeContainerProvider() *FakeContainerProvider {
	return &FakeContainerProvider{}
}

func (cp *FakeContainerProvider) GetContainer() (services.ContainerListModel, error) {

	addrs := make([]string, 0)
	ports := make([]string, 0)

	return services.ContainerListModel{
		ID:          "string",
		Name:        "string",
		Image:       "string",
		State:       "string",
		CreatedStr:  "string",
		IPAddresses: addrs,
		Ports:       ports,
	}, nil
}
