package tkui

import (
	"fmt"
	"hdu/internal/services"
	"hdu/internal/utils"
	"strings"

	"github.com/visualfc/atk/tk"
)

// область со списком контейнеров
type ImagesFrame struct {
	*tk.Frame
	tree        *tk.TreeView
	current_map map[string]services.ImageListModel
}

func NewImagesFrame(parent tk.Widget) *ImagesFrame {

	fr := tk.NewFrame(parent)

	// tree
	tree := tk.NewTreeView(fr)
	labels := []string{"REPOSITORY", "TAG", "IMAGE ID", "CREATED", "SIZE"}
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

	// instance
	current_map := make(map[string]services.ImageListModel)
	cf := &ImagesFrame{fr, tree, current_map}

	// events
	tree.BindEvent("<Double-1>", func(e *tk.Event) {
		items := tree.SelectionList()
		if len(items) > 0 {
			cf.onSelected(items[0])
		}
	})

	tree.BindEvent("<Return>", func(e *tk.Event) {
		items := tree.SelectionList()
		if len(items) > 0 {
			cf.onSelected(items[0])
		}
	})

	return cf

}

func (cf *ImagesFrame) SetItems(items []services.ImageListModel) {

	// clear all
	cf.tree.DeleteAllItems()
	for k := range cf.current_map {
		delete(cf.current_map, k)
	}

	// fill
	root := cf.tree.RootItem()
	for i, item := range items {

		for j, tag := range item.RepoTags {

			row_id := i*100 + j

			short_image_id := utils.ShortImageID(item.ID)
			display_size := utils.ByteCountToDisplaySize(item.Size)
			repo, short_tag := cf.splitTag(tag)

			row := root.InsertItem(row_id, repo, []string{short_tag, short_image_id, item.Created, display_size})
			// fmt.Println(row.Id())
			// fmt.Println(item.ID)
			cf.current_map[row.Id()] = item

		}
	}
}

func (cf *ImagesFrame) onSelected(item *tk.TreeItem) {

	utils.PrintAsJson(cf.current_map)

	fmt.Println(item.Id())
	model := cf.current_map[item.Id()]
	fmt.Println(model)

	// if cf.on_container_selected_handler != nil {
	// 	cf.on_container_selected_handler(model)
	// }
}

func (cf *ImagesFrame) splitTag(full_tag string) (string, string) {
	result := strings.Split(full_tag, ":")

	repo := strings.Join(result[:len(result)-1], ":")
	return repo, result[len(result)-1]
}
