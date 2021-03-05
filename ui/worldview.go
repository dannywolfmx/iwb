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
	actualChunk world.Chunk
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
	w.viewport.X += position
}

//TODO controle the unrange position. ej: -1
func (w *worldView) moveViewportY(position int) {
	w.viewport.Y += position
}

//TODO the printer dont works with special characters, just support 1 rune at the time
//TODO Pass the style by parameter
func (w *worldView) printOnScreen(text rune, viewport world.Position, wv, hv int) {
	//Print On Center of screen
	w.screen.SetContent(viewport.X-w.viewport.X+wv/2, viewport.Y-w.viewport.Y+hv/2, text, nil, tcell.StyleDefault)
	//Move the position to the next rune
}

func (w *worldView) Draw() {
	w.screen.Clear()
	wv, hv := w.screen.Size()
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
