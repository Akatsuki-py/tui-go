package main

import (
	"github.com/marcusolsson/tui-go"
)

func main() {
	var currentView int

	views := []tui.Widget{
		tui.NewVBox(tui.NewLabel("Press right arrow to continue ...")),
		tui.NewVBox(tui.NewLabel("Almost there, one more time!")),
		tui.NewVBox(tui.NewLabel("Congratulations, you've finished the example!")),
	}

	root := tui.NewVBox(views[0])

	ui := tui.New(root)
	ui.SetKeybinding(tui.KeyEsc, func() { ui.Quit() })
	ui.SetKeybinding(tui.KeyArrowLeft, func() {
		currentView = clamp(currentView-1, 0, len(views)-1)
		ui.SetWidget(views[currentView])
	})
	ui.SetKeybinding(tui.KeyArrowRight, func() {
		currentView = clamp(currentView+1, 0, len(views)-1)
		ui.SetWidget(views[currentView])
	})

	if err := ui.Run(); err != nil {
		panic(err)
	}
}

func clamp(n, min, max int) int {
	if n < min {
		return min
	}
	if n > max {
		return max
	}
	return n
}