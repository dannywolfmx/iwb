package ui

import (
	"fmt"
	"os"

	"github.com/dannywolfmx/iwb/world"
	"github.com/gdamore/tcell/v2"
)

type worldView struct {
	screen      tcell.Screen
	viewport    world.Position
	world       world.PersistantWorld
	actualChunk *world.Chunk
}

//NewWorldView create a worldView
func NewWorldView(screen tcell.Screen, w world.PersistantWorld) *worldView {
	return &worldView{
		screen:   screen,
		viewport: w.GetPosition(),
		world:    w,
		//TODO: Check the chunk sistem
		actualChunk: w.GetChunk(0, 0),
	}
}

//Clear will erase any character into the world screen
func (w *worldView) Clear() {
	w.screen.Clear()
}

//TODO controle the unrange position. ej: -1
func (w *worldView) moveViewportX(position int) {
	//If a use the position like uint i will lose the negative numbers, thats because i need to do this
	viewport := int(w.viewport.X) + position
	w.viewport.X = uint8(viewport)
}

//TODO controle the unrange position. ej: -1
func (w *worldView) moveViewportY(position int) {
	//If a use the position like uint i will lose the negative numbers, thats because i need to do this
	viewport := int(w.viewport.Y) + position
	w.viewport.Y = uint8(viewport)
}

//TODO the printer dont works with special characters, just support 1 rune at the time
//TODO Pass the style by parameter
//TODO Deal with the uint position: if i did 1 - 2 will be 255
func (w *worldView) printOnScreen(text rune, viewport world.Position, wv, hv int) {
	positionX := int(viewport.X) - int(w.viewport.X)
	positionY := int(viewport.Y) - int(w.viewport.Y)
	//Print On Center of screen
	w.screen.SetContent(positionX+wv/2, positionY+hv/2, text, nil, tcell.StyleDefault)
	//Move the position to the next rune
}

func (w *worldView) Draw() {
	w.screen.Clear()
	wv, hv := w.screen.Size()
	for viewport, text := range generateBorder() {
		w.printOnScreen(text, viewport, wv, hv)
	}
	for viewport, text := range w.actualChunk.GetElements() {
		w.printOnScreen(text, viewport, wv, hv)
	}
	w.screen.Show()
}

func (w *worldView) Run() {
	//Firs draw
	w.Draw()
	for {
		switch ev := w.screen.PollEvent().(type) {
		case *tcell.EventResize:
			w.Sync()
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyCtrlC:
				//CTRL + C to exit
				w.screen.Fini()
				w.world.SetPosition(w.viewport)
				if err := w.world.Persist(); err != nil {
					fmt.Println(err)
				}
				os.Exit(0)
			case tcell.KeyUp:
				//Move the viewPort to the UP
				w.moveViewportY(-1)
				w.Draw()
			case tcell.KeyDown:
				//Move the viewPort to the DOWN
				w.moveViewportY(1)
				w.Draw()
			case tcell.KeyRight:
				//Move the viewPort to the RIGHT
				w.moveViewportX(1)
				w.Draw()
			case tcell.KeyLeft:
				//Move the viewPort to the LEFT
				w.moveViewportX(-1)
				w.Draw()
			default:
				text := ev.Rune()
				w.actualChunk.SetElement(w.viewport, text)
				//Get Rune to print in the screen
				w.moveViewportX(1)
				w.Draw()
			}
		}
	}
}

//Sync will sync any single cell in the screen, it is more expensive than use tcell.Screen.Show()
func (w *worldView) Sync() {
	w.screen.Sync()
}

//generateBorder get a border of the actual chunk
//TODO: create a better border system using the actual viewport to eliminate unnecessary runes
func generateBorder() world.Elements {
	border := make(world.Elements)

	//TOP
	for i := 0; i <= 255; i++ {
		position := world.Position{X: uint8(i), Y: 0}
		border[position] = '-'
	}

	//BOTTOM
	for i := 0; i <= 255; i++ {
		position := world.Position{X: uint8(i), Y: 255}
		border[position] = '-'
	}

	//LEFT
	for i := 1; i <= 254; i++ {
		position := world.Position{X: 0, Y: uint8(i)}
		border[position] = '|'
	}

	//RIGHT
	for i := 1; i <= 254; i++ {
		position := world.Position{X: 255, Y: uint8(i)}
		border[position] = '|'
	}

	return border
}
