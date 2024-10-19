package tkui

import (
	"hdu/internal/services"

	"github.com/visualfc/atk/tk"
)

type ContainerViewProvider interface {
	GetContainer() (services.ContainerListModel, error)
}

type ContainerView struct {
	*tk.Frame
	provider ContainerViewProvider
}

func NewContainerView(parent tk.Widget, container_provider ContainerViewProvider) *ContainerView {

	fr := tk.NewFrame(parent)

	// label
	lbl := tk.NewLabel(fr, "Container")

	// tabs
	tabs := tk.NewNotebook(fr)

	container_info_page := NewContainerInfoPage(tabs)
	tabs.AddTab(container_info_page, "Main info")

	tabs.AddTab(tk.NewLabel(tabs, "logs"), "Logs")
	tabs.AddTab(tk.NewLabel(tabs, "stats"), "Stats")
	tabs.AddTab(tk.NewLabel(tabs, "env"), "Env")
	tabs.AddTab(tk.NewLabel(tabs, "config"), "Config")
	tabs.AddTab(tk.NewLabel(tabs, "top"), "Top")

	// layout
	main_layout := tk.NewVPackLayout(fr)
	main_layout.AddWidget(lbl)
	main_layout.AddWidget(tabs,
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
	)

	// fr.BindEvent("<Enter>", func(e *tk.Event) {
	// 	fmt.Println("Enter")
	// })
	//
	// fr.BindEvent("<Activate>", func(e *tk.Event) {
	// 	fmt.Println("Activate")
	// })

	page := &ContainerView{
		fr,
		container_provider,
	}

	// bind events --------------------------------------------------------------

	// событие отображения(при переключении вкладок)
	// fr.BindEvent("<Visibility>", func(e *tk.Event) {
	// 	page.refresh()
	// })
	//

	return page
}

// container info -------------------------------------------------------------
type ContainerInfoPage struct {
	*tk.Frame
}

func NewContainerInfoPage(root tk.Widget) *ContainerInfoPage {
	fr := tk.NewFrame(root)
	page := &ContainerInfoPage{fr}

	lbl := tk.NewLabel(fr, "Container info")

	// actions bar
	actions_bar := page.makeActionsBar(fr)
	status_bar := page.makeStatusBar(fr)
	details_bar := page.makeDetailsBar(fr)
	volumes_bar := page.makeVolumesBar(fr)
	networks_bar := page.makeNetworksBar(fr)

	// layout
	main_layout := tk.NewVPackLayout(fr)
	main_layout.AddWidget(lbl)
	main_layout.AddWidget(actions_bar, tk.PackAttrFillX())
	main_layout.AddWidget(status_bar, tk.PackAttrFillX())
	main_layout.AddWidget(details_bar, tk.PackAttrFillX())
	main_layout.AddWidget(volumes_bar, tk.PackAttrFillX())
	main_layout.AddWidget(networks_bar, tk.PackAttrFillX())

	return page
}

func (cip *ContainerInfoPage) makeActionsBar(root tk.Widget) *tk.LabelFrame {
	fr := tk.NewLabelFrame(root)
	fr.SetLabelText("Actions")

	btn_start := tk.NewButton(fr, "Start")
	btn_stop := tk.NewButton(fr, "Stop")
	btn_kill := tk.NewButton(fr, "Kill")
	btn_restart := tk.NewButton(fr, "Restart")
	btn_pause := tk.NewButton(fr, "Pause")
	btn_resume := tk.NewButton(fr, "Resume")
	btn_remove := tk.NewButton(fr, "Remove")

	// layout
	main_layout := tk.NewHPackLayout(fr)
	main_layout.AddWidget(btn_start)
	main_layout.AddWidget(btn_stop)
	main_layout.AddWidget(btn_kill)
	main_layout.AddWidget(btn_restart)
	main_layout.AddWidget(btn_pause)
	main_layout.AddWidget(btn_resume)
	main_layout.AddWidget(btn_remove)

	return fr
}

func (cip *ContainerInfoPage) makeStatusBar(root tk.Widget) *tk.LabelFrame {
	fr := tk.NewLabelFrame(root)
	fr.SetLabelText("Status")

	btn_start := tk.NewButton(fr, "Start")
	btn_stop := tk.NewButton(fr, "Stop")
	btn_kill := tk.NewButton(fr, "Kill")
	btn_restart := tk.NewButton(fr, "Restart")
	btn_pause := tk.NewButton(fr, "Pause")
	btn_resume := tk.NewButton(fr, "Resume")
	btn_remove := tk.NewButton(fr, "Remove")

	// layout
	main_layout := tk.NewHPackLayout(fr)
	main_layout.AddWidget(btn_start)
	main_layout.AddWidget(btn_stop)
	main_layout.AddWidget(btn_kill)
	main_layout.AddWidget(btn_restart)
	main_layout.AddWidget(btn_pause)
	main_layout.AddWidget(btn_resume)
	main_layout.AddWidget(btn_remove)

	return fr
}

func (cip *ContainerInfoPage) makeDetailsBar(root tk.Widget) *tk.LabelFrame {
	fr := tk.NewLabelFrame(root)
	fr.SetLabelText("Details")

	btn_start := tk.NewButton(fr, "Start")
	btn_stop := tk.NewButton(fr, "Stop")
	btn_kill := tk.NewButton(fr, "Kill")
	btn_restart := tk.NewButton(fr, "Restart")
	btn_pause := tk.NewButton(fr, "Pause")
	btn_resume := tk.NewButton(fr, "Resume")
	btn_remove := tk.NewButton(fr, "Remove")

	// layout
	main_layout := tk.NewHPackLayout(fr)
	main_layout.AddWidget(btn_start)
	main_layout.AddWidget(btn_stop)
	main_layout.AddWidget(btn_kill)
	main_layout.AddWidget(btn_restart)
	main_layout.AddWidget(btn_pause)
	main_layout.AddWidget(btn_resume)
	main_layout.AddWidget(btn_remove)

	return fr
}

func (cip *ContainerInfoPage) makeVolumesBar(root tk.Widget) *tk.LabelFrame {
	fr := tk.NewLabelFrame(root)
	fr.SetLabelText("Volumes")

	btn_start := tk.NewButton(fr, "Start")
	btn_stop := tk.NewButton(fr, "Stop")
	btn_kill := tk.NewButton(fr, "Kill")
	btn_restart := tk.NewButton(fr, "Restart")
	btn_pause := tk.NewButton(fr, "Pause")
	btn_resume := tk.NewButton(fr, "Resume")
	btn_remove := tk.NewButton(fr, "Remove")

	// layout
	main_layout := tk.NewHPackLayout(fr)
	main_layout.AddWidget(btn_start)
	main_layout.AddWidget(btn_stop)
	main_layout.AddWidget(btn_kill)
	main_layout.AddWidget(btn_restart)
	main_layout.AddWidget(btn_pause)
	main_layout.AddWidget(btn_resume)
	main_layout.AddWidget(btn_remove)

	return fr
}

func (cip *ContainerInfoPage) makeNetworksBar(root tk.Widget) *tk.LabelFrame {
	fr := tk.NewLabelFrame(root)
	fr.SetLabelText("Networks")

	btn_start := tk.NewButton(fr, "Start")
	btn_stop := tk.NewButton(fr, "Stop")
	btn_kill := tk.NewButton(fr, "Kill")
	btn_restart := tk.NewButton(fr, "Restart")
	btn_pause := tk.NewButton(fr, "Pause")
	btn_resume := tk.NewButton(fr, "Resume")
	btn_remove := tk.NewButton(fr, "Remove")

	// layout
	main_layout := tk.NewHPackLayout(fr)
	main_layout.AddWidget(btn_start)
	main_layout.AddWidget(btn_stop)
	main_layout.AddWidget(btn_kill)
	main_layout.AddWidget(btn_restart)
	main_layout.AddWidget(btn_pause)
	main_layout.AddWidget(btn_resume)
	main_layout.AddWidget(btn_remove)

	return fr
}
