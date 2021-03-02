package ui

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

//Alis to better reading
type position = int

type worldView struct {
	screen                        tcell.Screen
	cursorX, cursorY              position
	viewportX, viewportY          position
	viewportWidth, viewportHeight position
	needUpdate                    bool
}

//NewWorldView create a worldView
func NewWorldView(screen tcell.Screen) *worldView {

	return &worldView{
		screen: screen,
	}
}

//Clear will erase any character into the world screen
func (w *worldView) Clear() {
	w.screen.Clear()
}

//TODO controle the unrange position. ej: -1
func (w *worldView) moveViewportX(newPosition position) {
	w.viewportX += newPosition
}

//TODO controle the unrange position. ej: -1
func (w *worldView) moveViewportY(newPosition position) {
	w.viewportY += newPosition
}

//TODO the printer dont works with special characters, just support 1 rune at the time
//TODO Pass the style by parameter
func (w *worldView) printOnScreen(text rune) {
	w.screen.SetContent(w.viewportX, w.viewportY, text, nil, tcell.StyleDefault)
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
