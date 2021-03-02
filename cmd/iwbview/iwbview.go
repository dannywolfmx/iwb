package main

import (
	"os"

	"github.com/dannywolfmx/iwb/ui"
	"github.com/dannywolfmx/iwb/world/memory"
	"github.com/gdamore/tcell/v2"
)

func main() {

	world := memory.NewMemoryWorld()
	worldUI := ui.NewWorldView(world)

	for {
		switch ev := worldUI.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			worldUI.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				worldUI.Screen.Fini()
				os.Exit(0)
			}
		}

	}
}
