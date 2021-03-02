package main

import (
	"os"

	"github.com/dannywolfmx/iwb/ui"
)

func main() {
	style := ui.DefaultStyle()
	screen, err := ui.NewDefaultScreen(style)

	if err != nil {
		os.Exit(1)
	}
	ui.NewWorldView(screen).Run()
}
