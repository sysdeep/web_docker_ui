package containers_page

import (
	"fmt"
	"hdu/internal/services"
	"sort"

	"github.com/visualfc/atk/tk"
)

// type onContainerSelectedHandler = func(model *services.ContainerListModel)

// область со списком контейнеров
type listBar struct {
	*tk.Frame
	tree          *tk.TreeView
	current_map   map[string]services.ContainerListModel
	vm            *ContainersVM
	current_items []services.ContainerListModel
}

func newListBar(parent tk.Widget, vm *ContainersVM) *listBar {

	fr := tk.NewFrame(parent)

	// tree
	tree := tk.NewTreeView(fr)
	labels := []string{"state", "name", "image"}
	tree.SetColumnCount(len(labels))
	tree.SetHeaderLabels(labels)

	// help
	help_frame := tk.NewFrame(fr)
	help_label_delete := tk.NewLabel(help_frame, "delete: [d]")
	help_label_stop := tk.NewLabel(help_frame, "stop: [s]")
	help_layout := tk.NewHPackLayout(help_frame)
	help_layout.AddWidgets(help_label_delete, help_label_stop)

	// layout
	main_layout := tk.NewVPackLayout(fr)
	main_layout.AddWidget(tree,
		tk.PackAttrFillBoth(),
		tk.PackAttrPadx(4),
		tk.PackAttrPady(4),
		tk.PackAttrExpand(true),
	)
	main_layout.AddWidget(help_frame,
		tk.PackAttrFillX(),
	)

	current_map := make(map[string]services.ContainerListModel)
	cf := &listBar{
		Frame:         fr,
		tree:          tree,
		current_map:   current_map,
		vm:            vm,
		current_items: make([]services.ContainerListModel, 0),
		// on_container_selected_handler: nil,
	}

	vm.register(cf)

	// events
	tree.BindEvent("<Double-1>", func(e *tk.Event) {
		items := tree.SelectionList()
		// fmt.Println(items)
		if len(items) > 0 {
			cf.onSelected(items[0])
		}
	})

	tree.BindEvent("<Return>", func(e *tk.Event) {
		items := tree.SelectionList()
		// fmt.Println(items)
		if len(items) > 0 {
			cf.onSelected(items[0])
		}
	})

	// bind d for delete
	tree.BindEvent("<KeyRelease-d>", func(e *tk.Event) {
		fmt.Println("d pressed")

		items := tree.SelectionList()
		if len(items) > 0 {
			model := cf.current_map[items[0].Id()]
			fmt.Println("do delete: " + model.Name)
		}

	})

	// bind s for stop
	tree.BindEvent("<KeyRelease-s>", func(e *tk.Event) {
		fmt.Println("s pressed")

		items := tree.SelectionList()
		if len(items) > 0 {
			model := cf.current_map[items[0].Id()]
			fmt.Println("do stop: " + model.Name)
		}

	})

	return cf
}

// func (cf *ContainersFrame) SetItems(items []services.ContainerListModel) {
//
// 	// clear all
// 	cf.tree.DeleteAllItems()
// 	for k := range cf.current_map {
// 		delete(cf.current_map, k)
// 	}
//
// 	// fill
// 	root := cf.tree.RootItem()
// 	for i, item := range items {
// 		row := root.InsertItem(i*10, item.State, []string{item.Name, item.Image})
// 		fmt.Println(row.Id())
// 		cf.current_map[row.Id()] = item
// 	}
// }

func (cf *listBar) onSelected(item *tk.TreeItem) {
	// fmt.Println(item)
	// fmt.Println(item.Id(), item.Index(), item.Values())

	model := cf.current_map[item.Id()]
	// utils.PrintAsJson(model)

	// if cf.on_container_selected_handler != nil {
	// 	cf.on_container_selected_handler(&model)
	// }

	cf.vm.show_container(model)
}

// func (cf *ContainersFrame) ConnectContainerSelected(handler onContainerSelectedHandler) {
// 	cf.on_container_selected_handler = handler
// }

func (cf *listBar) update(ev string) {
	switch ev {
	case EV_REFRESH:
		cf.refresh()
		break
	case EV_FILTER:
		cf.fill()
		break

	}
}

func (cf *listBar) refresh() {

	cf.current_items = cf.vm.get_containers()
	cf.fill()
}

func (cf *listBar) fill() {

	var items []services.ContainerListModel

	current_filter := cf.vm.get_filter()

	// filter status
	for _, cnt := range cf.current_items {
		if current_filter.status == FILTER_ST_ALL {
			items = append(items, cnt)
		}

		if current_filter.status == FILTER_ST_RUNNING && cnt.State == "running" {
			items = append(items, cnt)
		}

		if current_filter.status == FILTER_ST_STOPPED && cnt.State == "stopped" {
			items = append(items, cnt)
		}
	}

	// sort
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Name < items[j].Name
	})

	// clear all
	cf.tree.DeleteAllItems()
	for k := range cf.current_map {
		delete(cf.current_map, k)
	}

	// fill
	root := cf.tree.RootItem()
	for i, item := range items {
		row := root.InsertItem(i*10, item.State, []string{item.Name, item.Image})
		fmt.Println(row.Id())

		cf.current_map[row.Id()] = item
	}

}
