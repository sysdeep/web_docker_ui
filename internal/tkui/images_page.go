package tkui

import (
	"fmt"
	"hdu/internal/services"

	"github.com/visualfc/atk/tk"
)

type ImagesPage struct {
	*tk.Frame
	images_service *services.ImagesService
	list_frame     *ImagesFrame
}

func NewImagesPage(parent tk.Widget, images_service *services.ImagesService) *ImagesPage {

	page := &ImagesPage{
		Frame:          tk.NewFrame(parent),
		images_service: images_service,
	}

	// label
	lbl := tk.NewLabel(page.Frame, "Images")

	// list
	page.list_frame = NewImagesFrame(page.Frame)

	// actions bar
	actions_bar := page._make_controls_bar(page.Frame)

	// lauout
	main_layout := tk.NewVPackLayout(page.Frame)
	main_layout.AddWidget(lbl)
	main_layout.AddWidget(page.list_frame,
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
	)
	main_layout.AddWidget(actions_bar,
		tk.PackAttrFillX(),
		// tk.PackAttrExpand(true),
	)

	// fr.BindEvent("<Enter>", func(e *tk.Event) {
	// 	fmt.Println("Enter")
	// })
	//
	// fr.BindEvent("<Activate>", func(e *tk.Event) {
	// 	fmt.Println("Activate")
	// })

	// page := &ImagesPage{
	// 	fr,
	// 	images_service,
	// 	list,
	// }

	// bind events --------------------------------------------------------------

	// событие отображения(при переключении вкладок)
	page.Frame.BindEvent("<Visibility>", func(e *tk.Event) {
		page.refresh()
	})

	return page
}

func (cp *ImagesPage) refresh() {
	fmt.Println("images refresh")
	items, _ := cp.images_service.GetAll()

	cp.list_frame.SetItems(items)
}

func (cp *ImagesPage) _make_controls_bar(parent tk.Widget) *tk.Frame {
	fr := tk.NewFrame(parent)

	btn_refresh := tk.NewButton(fr, "refresh")

	layout := tk.NewHPackLayout(fr)

	layout.AddWidget(tk.NewLayoutSpacer(fr, 1, true))
	layout.AddWidget(btn_refresh,
		tk.PackAttrSideRight(),
	)

	// events -------------------------------------------------------------------
	btn_refresh.OnCommand(func() {
		cp.refresh()
	})

	return fr
}
