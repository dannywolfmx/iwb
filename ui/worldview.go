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
func (w *worldView) printOnScreen(text rune) {
	w.state[w.viewport] = text
	w.screen.SetContent(w.viewport.X, w.viewport.Y, text, nil, tcell.StyleDefault)
	//Move the position to the next rune
	w.moveViewportX(1)
	w.needUpdate = true
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
			case tcell.KeyDown:
				//Move the viewPort to the DOWN
				w.moveViewportY(1)
			case tcell.KeyRight:
				//Move the viewPort to the RIGHT
				w.moveViewportX(1)
			case tcell.KeyLeft:
				//Move the viewPort to the LEFT
				w.moveViewportX(-1)
			default:
				//Get Rune to print in the screen
				w.printOnScreen(ev.Rune())
			}
		}
		//Check if the map need to repaint
		if w.needUpdate {
			w.screen.Show()
			w.needUpdate = false
		}
	}
}

//Sync will sync any single cell in the screen, it is more expensive than use tcell.Screen.Show()
func (w *worldView) Sync() {
	w.screen.Sync()
}
