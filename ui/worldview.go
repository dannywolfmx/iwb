package ui

import (
	"github.com/dannywolfmx/iwb/world"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type worldView struct {
	*tview.Grid
	App                           *tview.Application
	cursorX, cursorY              int
	viewportX, viewportY          int
	viewportWidth, viewportHeight int
	world                         world.World
}

//NewWorldView create a worldView
func NewWorldView(world world.World, app *tview.Application) *worldView {
	view := &worldView{
		Grid:  newDefaultGrid(app),
		App:   app,
		world: world,
	}
	SetCaptureInput(view)

	return view
}
func newPrimitive(color tcell.Color) tview.Primitive {
	return tview.NewTextView().SetBackgroundColor(color)
}

//newDefaultGrid will set a default data to tview.Grid struct
//
//* Show a border arround the terminal
//* Show a chess like board
func newDefaultGrid(app *tview.Application) *tview.Grid {
	rowNum := 20
	collNum := 20
	grid := tview.NewGrid().SetSize(rowNum, collNum, 0, 0)
	grid.SetBorder(true)

	return grid
}

//fillChesBoard need to be runned with gorutine
//EJ... go fillChessBoard(app, grid, rowNum, collNum)
func fillChessBoard(app *tview.Application, grid *tview.Grid, rowNum int, collNum int) []tview.Primitive {
	cell := []tview.Primitive{}
	drawCell := true

	app.QueueUpdateDraw(func() {
		for row := 0; row <= rowNum-1; row++ {
			for coll := 0; coll <= collNum; coll++ {
				if drawCell = !drawCell; drawCell {
					primitive := newPrimitive(tcell.Color100)
					cell = append(cell, primitive)
					grid.AddItem(primitive, row, coll, 1, 1, 0, 0, true)
				}
			}
		}
	})
	return cell
}

func SetCaptureInput(view *worldView) {
	view.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Key() == tcell.KeyDown {
			view.moveElement(1, 0)
		} else if event.Key() == tcell.KeyUp {
			view.moveElement(-1, 0)
		} else if event.Key() == tcell.KeyLeft {
			view.moveElement(0, -1)
		} else if event.Key() == tcell.KeyRight {
			view.moveElement(0, 1)
		} else if event.Key() == tcell.KeyRune {
		} else {
			return event
		}
		return nil
	})
}

func (v *worldView) moveElement(x, y int) {
	v.cursorX += x
	v.cursorY += y
	if v.cursorX < 0 || v.cursorY < 0 {
		//Revert position
		v.cursorX -= x
		v.cursorY -= y
	}

	v.Grid.Clear()
	primitive := newPrimitive(tcell.Color200)
	v.AddItem(primitive, v.cursorX, v.cursorY, 1, 1, 1, 1, true)
}
