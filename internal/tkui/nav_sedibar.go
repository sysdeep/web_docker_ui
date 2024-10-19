package tkui

import "github.com/visualfc/atk/tk"

type NavSidebar struct {
	*tk.Frame
	tree *tk.TreeView
}

func NewNavSidebar(parent tk.Widget) *NavSidebar {
	sb := &NavSidebar{}
	sb.Frame = tk.NewFrame(parent)

	sb.tree = tk.NewTreeView(sb)
	sb.tree.SetColumnCount(1)
	sb.tree.SetHeaderHidden(true)

	lbl := tk.NewLabel(sb, "sidebar")

	// layout
	vbox := tk.NewVPackLayout(sb)
	vbox.AddWidget(lbl)
	vbox.AddWidget(sb.tree, tk.PackAttrFillBoth(), tk.PackAttrExpand(true))

	// init
	sb.fillTree()

	return sb
}

func (sb *NavSidebar) fillTree() {
	root := sb.tree.RootItem()
	docker := sb.tree.InsertItem(root, 1, "docker", []string{})
	sb.tree.InsertItem(docker, 2, "containers", []string{})
	sb.tree.InsertItem(root, 101, "swarm", []string{})
	sb.tree.InsertItem(root, 201, "registry", []string{})

}
