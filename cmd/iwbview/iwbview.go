package main

import (
	"github.com/dannywolfmx/iwb/ui"
	"github.com/dannywolfmx/iwb/world/memory"
)

func main() {

	world := memory.NewMemoryWorld()
	ui.NewWorldView(world).Run()
}
