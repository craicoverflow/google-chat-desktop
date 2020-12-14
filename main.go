package main

import (
	"github.com/zserge/lorca"
)

func main() {
	ui, _ := lorca.New("https://chat.google.com", "", 1000, 800)
	defer ui.Close()

	// Wait for the browser window to be closed
	<-ui.Done()
}
