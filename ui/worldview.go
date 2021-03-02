package ui

import (
	"os"

	"github.com/dannywolfmx/iwb/world"
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/rivo/tview"
)

type worldView struct {
	screen                        tcell.Screen
	cursorX, cursorY              int
	viewportX, viewportY          int
	viewportWidth, viewportHeight int
	world                         world.World
}

func (w *worldView) Sync() {
	w.screen.Sync()
}

func (w *worldView) Run() {
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

//NewWorldView create a worldView
func NewWorldView(world world.World) *worldView {

	encoding.Register()

	style := DefaultStyle()
	screen, err := newDefaultScreen(style)

	if err != nil {
		return nil
	}

	view := &worldView{
		screen: screen,
		world:  world,
	}
	SetCaptureInput(view)

	return view
}

func DefaultStyle() tcell.Style {
	return tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)
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

func SetCaptureInput(view *worldView) {
	//	view.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	//
	//		if event.Key() == tcell.KeyDown {
	//			view.moveElement(1, 0)
	//		} else if event.Key() == tcell.KeyUp {
	//			view.moveElement(-1, 0)
	//		} else if event.Key() == tcell.KeyLeft {
	//			view.moveElement(0, -1)
	//		} else if event.Key() == tcell.KeyRight {
	//			view.moveElement(0, 1)
	//		} else if event.Key() == tcell.KeyRune {
	//		} else {
	//			return event
	//		}
	//		return nil
	//	})
}

func (v *worldView) moveElement(x, y int) {
	//	v.cursorX += x
	//	v.cursorY += y
	//	if v.cursorX < 0 || v.cursorY < 0 {
	//		//Revert position
	//		v.cursorX -= x
	//		v.cursorY -= y
	//	}
	//
	//	v.Grid.Clear()
	//	primitive := newPrimitive(tcell.Color200)
	//	v.AddItem(primitive, v.cursorX, v.cursorY, 1, 1, 1, 1, true)
}
