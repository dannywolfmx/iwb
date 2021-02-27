package main

import (
	"github.com/dannywolfmx/iwb/ui"
	"github.com/dannywolfmx/iwb/world/memory"
	"github.com/rivo/tview"
)

func main() {
	world := memory.NewMemoryWorld()
	box := ui.NewWorldView(world)
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
