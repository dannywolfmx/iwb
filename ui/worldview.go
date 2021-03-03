package ui

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

type Position struct {
	X, Y int
}

type worldView struct {
	screen                        tcell.Screen
	cursorX, cursorY              int
	viewportWidth, viewportHeight int
	needUpdate                    bool
	viewport                      Position
	state                         map[Position]rune
}

//NewWorldView create a worldView
func NewWorldView(screen tcell.Screen) *worldView {
	return &worldView{
		screen: screen,
		state:  make(map[Position]rune),
	}
}

//Clear will erase any character into the world screen
func (w *worldView) Clear() {
	w.screen.Clear()
}

//TODO controle the unrange position. ej: -1
func (w *worldView) moveViewportX(newPosition int) {
	w.viewport.X += newPosition
}

//TODO controle the unrange position. ej: -1
func (w *worldView) moveViewportY(newPosition int) {
	w.viewport.Y += newPosition
}

//TODO the printer dont works with special characters, just support 1 rune at the time
//TODO Pass the style by parameter
func (w *worldView) printOnScreen(text rune, viewport Position, wv, hv int) {
	//Print On Center of screen
	w.screen.SetContent(viewport.X-w.viewport.X+wv/2, viewport.Y-w.viewport.Y+hv/2, text, nil, tcell.StyleDefault)
	//Move the position to the next rune
}

func (w *worldView) Draw() {
	w.screen.Clear()
	wv, hv := w.screen.Size()
	for viewport, text := range w.state {
		w.printOnScreen(text, viewport, wv, hv)
	}
	w.screen.Show()
}

func (w *worldView) Run() {
	for {
		switch ev := w.screen.PollEvent().(type) {
		case *tcell.EventResize:
			w.Sync()
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyCtrlC:
				//CTRL + C to exit
				w.screen.Fini()
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
				w.state[w.viewport] = text
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
