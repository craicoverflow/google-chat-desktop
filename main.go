package main

import (
	"os"

	"github.com/webview/webview"
)

func main() {
	debug := os.Getenv("DEBUG") == "true"
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Google Chat Desktop (Unofficial)")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://chat.google.com")
	w.Run()
}
