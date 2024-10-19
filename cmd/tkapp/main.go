package main

import (
	"fmt"
	"hdu/internal/services"
	"hdu/internal/tkui"

	"github.com/docker/docker/client"
	"github.com/visualfc/atk/tk"
)

type Window struct {
	*tk.Window
}

func main() {

	// docker client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// init services
	servs := services.NewServices(cli)

	// start mainloop
	tk.MainLoop(func() {
		mw := tkui.NewMainWindow(servs)
		// mw.ResizeN(1024, 800)
		// mw := NewWindow()
		mw.SetTitle("Docker Sample")
		// mw.Center(nil)
		mw.ShowNormal()
		// mw.ShowMaximized()
		// mw.ShowFullScreen()
		fmt.Println(tk.DumpWidget(mw))
	})

}
