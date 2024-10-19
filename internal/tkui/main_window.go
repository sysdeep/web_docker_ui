package tkui

import (
	"fmt"
	"hdu/internal/services"
	"hdu/internal/tkui/containers_page"
	"hdu/internal/tkui/volumes_page"

	"github.com/visualfc/atk/tk"
)

type MainWindow struct {
	*tk.Window
	// containers_page *ContainersPage
}

func NewMainWindow(servs *services.Services) *MainWindow {

	mw := &MainWindow{
		tk.RootWindow(),
	}

	// content
	content_frame := tk.NewFrame(mw)
	content_layout := tk.NewHPackLayout(content_frame)

	// sidebar := NewNavSidebar(mw)
	tabs := mw.makeTabs(mw, servs)
	actions_bar := mw.makeActionsBar(mw)

	// layout
	// content_layout.AddWidget(sidebar, tk.PackAttrFillY(), tk.PackAttrSideLeft())
	content_layout.AddWidget(tabs,
		// tk.PackAttrFillY(),
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
		tk.PackAttrSideRight())

	main_layout := tk.NewVPackLayout(mw)
	main_layout.AddWidget(content_frame,
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
		tk.PackAttrPadx(4),
		tk.PackAttrPady(4),
	)
	main_layout.AddWidget(actions_bar,
		tk.PackAttrSideBottom(),
		tk.PackAttrExpand(false),
		tk.PackAttrFillX(),
		tk.PackAttrPadx(4),
		tk.PackAttrPady(4),
	)

	return mw

}

func (mw *MainWindow) makeTabs(root tk.Widget, servs *services.Services) *tk.Notebook {
	tabs := tk.NewNotebook(root)

	actions_handler := &RootActionsHandler{}

	containers_page := containers_page.NewContainersPage(tabs, containers_page.NewContainersVM(servs.Containers, actions_handler))
	tabs.AddTab(containers_page, "Containers")

	images_page := NewImagesPage(tabs, servs.Images)
	tabs.AddTab(images_page, "Images")

	volumes_page := volumes_page.NewVolumesPage(tabs, volumes_page.NewVolumesPageVM(servs.Volumes))
	tabs.AddTab(volumes_page, "Volumes")

	test_label := tk.NewLabel(tabs, "test")
	tabs.AddTab(test_label, "Test")

	return tabs
}

func (mw *MainWindow) makeActionsBar(root tk.Widget) *tk.Frame {
	fr := tk.NewFrame(root)

	exit_btn := tk.NewButton(mw, "Exit")
	exit_btn.OnCommand(mw.exit)

	layout := tk.NewHPackLayout(fr)
	layout.AddWidget(tk.NewLayoutSpacer(fr, 1, true))
	layout.AddWidget(exit_btn)

	return fr
}

func (mw *MainWindow) exit() {

	tk.Quit()
}

// TODO: move to file
type RootActionsHandler struct{}

func (ah *RootActionsHandler) ShowContainer(container_id string) {
	fmt.Println("root: show container")

	// TODO: in new type
	top := tk.NewWindow()

	// view := NewContainerView(top, NewFakeContainerProvider())
	// layout := tk.NewVPackLayout(top)
	// layout.AddWidget(view, tk.PackAttrFillBoth(), tk.PackAttrExpand(true))

	top.SetTitle("Container view")
	top.ShowNormal()

}
