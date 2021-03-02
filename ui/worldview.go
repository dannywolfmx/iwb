package ui

import (
	"os"

	"github.com/dannywolfmx/iwb/world"
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/rivo/tview"
)

//Alis to better reading
type position = int

type worldView struct {
	screen                        tcell.Screen
	cursorX, cursorY              position
	viewportX, viewportY          position
	viewportWidth, viewportHeight position
	world                         world.World
	needUpdate                    bool
}

//Sync will sync any single cell in the screen, it is more expensive than use tcell.Screen.Show()
func (w *worldView) Sync() {
	w.screen.Sync()
}

//Clear will erase any character into the world screen
func (w *worldView) Clear() {
	w.screen.Clear()
}

//TODO controle the unrange position. ej: -1
func (w *worldView) moveViewPortX(newPosition position) {
	w.viewportX += newPosition
}

//TODO controle the unrange position. ej: -1
func (w *worldView) moveViewPortY(newPosition position) {
	w.viewportY += newPosition
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
				w.moveViewPortY(-1)
			case tcell.KeyDown:
				//Move the viewPort to the DOWN
				w.moveViewPortY(1)
			case tcell.KeyRight:
				//Move the viewPort to the RIGHT
				w.moveViewPortX(1)
			case tcell.KeyLeft:
				//Move the viewPort to the LEFT
				w.moveViewPortX(-1)
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

//TODO the printer dont works with special characters, just support 1 rune at the time
//TODO Pass the style by parameter
func (w *worldView) printOnScreen(text rune) {
	w.screen.SetContent(w.viewportX, w.viewportY, text, nil, tcell.StyleDefault)
	//Move the position to the next rune
	w.moveViewPortX(1)
	w.needUpdate = true
}

//NewWorldView create a worldView
func NewWorldView(world world.World) *worldView {

	encoding.Register()

	style := DefaultStyle()
	screen, err := newDefaultScreen(style)

	if err != nil {
		return nil
	}

	return &worldView{
		screen: screen,
		world:  world,
	}
}

func DefaultStyle() tcell.Style {
	return tcell.StyleDefault
}

func newPrimitive(color tcell.Color) tview.Primitive {
	return tview.NewTextView().SetBackgroundColor(color)
}

//newDefaultGrid will set a default data to tview.Grid struct
//
//* Show a border arround the terminal
//* Show a chess like board
func newDefaultScreen(style tcell.Style) (tcell.Screen, error) {
	screen, err := tcell.NewScreen()

	if err != nil {
		return screen, err
	}

	if err = screen.Init(); err != nil {
		return screen, err
	}

	screen.SetStyle(style)

	return screen, nil
}
