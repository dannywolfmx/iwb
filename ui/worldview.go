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
}

//Sync will sync any single cell in the screen, it is more expensive than use tcell.Screen.Show()
func (w *worldView) Sync() {
	w.screen.Sync()
}

//Clear will erase any character into the world screen
func (w *worldView) Clear() {
	w.screen.Clear()
}

func (w *worldView) Run() {
	w.printOnScreen(1, 1, "Prueba")
	w.printOnScreen(1, 2, "Prueba")
	for {
		switch ev := w.screen.PollEvent().(type) {
		case *tcell.EventResize:
			w.Sync()
		case *tcell.EventKey:
			eventKey := ev.Key()
			if eventKey == tcell.KeyEscape || eventKey == tcell.KeyCtrlC {
				w.screen.Fini()
				os.Exit(0)
			}
		}

	}
}

//TODO the printer dont works with special characters, just support 1 rune at the time
//TODO Pass the style by parameter
func (w *worldView) printOnScreen(x, y position, text string) {
	for _, c := range text {
		comb := []rune{}
		w.screen.SetContent(x, y, c, comb, tcell.StyleDefault)
		//Move the position to the next rune
		x++
	}
	w.screen.Show()

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
