package volumes_page

import (
	"hdu/internal/services"
	"sort"

	"github.com/visualfc/atk/tk"
)

type listBar struct {
	*tk.Frame
	tree        *tk.TreeView
	vm          *VolumesPageVM
	current_map map[string]services.VolumeListModel
}

func newListBar(parent tk.Widget, vm *VolumesPageVM) *listBar {
	fr := tk.NewFrame(parent)
	// tree
	tree := tk.NewTreeView(fr)
	labels := []string{"NAME", "MOUNT", "STACK", "CREATED"}
	tree.SetColumnCount(len(labels))
	tree.SetHeaderLabels(labels)

	// layout
	main_layout := tk.NewVPackLayout(fr)
	main_layout.AddWidget(tree,
		tk.PackAttrFillBoth(),
		tk.PackAttrPadx(4),
		tk.PackAttrPady(4),
		tk.PackAttrExpand(true),
	)

	bar := &listBar{
		Frame:       fr,
		tree:        tree,
		vm:          vm,
		current_map: make(map[string]services.VolumeListModel),
	}

	vm.register(bar)
	return bar

}

func (b *listBar) update(ev string) {
	switch ev {
	case EV_REFRESH:
		b.refresh()
		break

	}
}

func (b *listBar) refresh() {

	volumes := b.vm.get_volumes()

	sort.SliceStable(volumes, func(i, j int) bool {
		return volumes[i].Name < volumes[j].Name
	})

	// clear all
	b.tree.DeleteAllItems()
	for k := range b.current_map {
		delete(b.current_map, k)
	}

	// fill
	root := b.tree.RootItem()
	for i, item := range volumes {

		row_id := i * 100

		row := root.InsertItem(row_id, item.Name, []string{item.Mountpoint, item.StackName, item.CreatedAt})
		b.current_map[row.Id()] = item

	}

}
