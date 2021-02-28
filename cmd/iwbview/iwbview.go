package main

import (
	"github.com/dannywolfmx/iwb/ui"
	"github.com/dannywolfmx/iwb/world/memory"
	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()

	world := memory.NewMemoryWorld()
	worldView := ui.NewWorldView(world, app)

	if err := app.SetRoot(worldView, true).Run(); err != nil {
		panic(err)
	}
}
