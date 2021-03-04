package main

import (
	"os"

	"github.com/dannywolfmx/iwb/ui"
	"github.com/dannywolfmx/iwb/world/memory"
)

func main() {
	style := ui.DefaultStyle()
	screen, err := ui.NewDefaultScreen(style)
	world := memory.NewMemoryWorld()

	if err != nil {
		os.Exit(1)
	}
	ui.NewWorldView(screen, world).Run()
}
