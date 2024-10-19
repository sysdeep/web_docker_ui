package main

import webview "github.com/webview/webview_go"

/*
sudo apt install libwebkit2gtk-4.0-dev
*/
func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(480, 320, webview.HintNone)
	// w.SetHtml("Thanks for using webview!")
	// w.Navigate("https://ya.ru")
	w.Navigate("http://localhost:1313")
	w.Run()
}
