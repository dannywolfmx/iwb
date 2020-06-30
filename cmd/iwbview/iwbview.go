package main

import (
	"danirod.es/pkg/iwb/ui"
	"danirod.es/pkg/iwb/world/memory"
	"github.com/rivo/tview"
)

func main() {
	world := memory.NewMemoryWorld()
	box := ui.NewWorldView(world)
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
