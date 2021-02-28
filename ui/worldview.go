package ui

import (
	"fmt"

	"github.com/dannywolfmx/iwb/world"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type worldView struct {
	*tview.Grid
	cursorX, cursorY              int
	viewportX, viewportY          int
	viewportWidth, viewportHeight int
	world                         world.World
}

//NewWorldView create a worldView
func NewWorldView(world world.World) *worldView {
	view := &worldView{
		Grid:  newDefaultGrid(),
		world: world,
	}

	return view
}
func newPrimitive(text string) tview.Primitive {
	return tview.NewTextView().SetText(text).SetTextColor(tcell.Color115).SetBackgroundColor(tcell.Color100)
}

type Grid struct {
}

//newDefaultGrid will set a default data to tview.Grid struct
//
//* Show a border arround the terminal
//* Show a chess like board
func newDefaultGrid() *tview.Grid {
	rowNum := 254
	collNum := 254
	grid := tview.NewGrid().SetSize(rowNum, collNum, 1, 1)

	cell := []tview.Primitive{}

	drawCell := true
	for row := 0; row <= rowNum-1; row++ {
		for coll := 0; coll <= collNum; coll++ {
			if drawCell = !drawCell; drawCell {
				primitive := newPrimitive(fmt.Sprintf("%d,%d", row, coll))
				cell = append(cell, primitive)
				grid.AddItem(primitive, row, coll, 1, 1, 1, 1, true)
			}
		}
	}
	return grid
}

func (w *worldView) moveCursorUp() {
	if w.cursorY == 0 {
		return
	}
	w.cursorY--
	if w.viewportY > 0 && w.cursorY < w.viewportY {
		w.viewportY--
	}
}

func (w *worldView) moveCursorDown() {
	if w.cursorY == (256*256 - 1) {
		return
	}
	w.cursorY++
	if w.viewportY < (256*256-w.viewportHeight-1) && w.cursorY >= w.viewportY+w.viewportHeight {
		w.viewportY++
	}
}

func (w *worldView) moveCursorLeft() {
	if w.cursorX == 0 {
		return
	}
	w.cursorX--
	if w.viewportX >= 0 && w.cursorX < w.viewportX {
		w.viewportX--
	}
}

func (w *worldView) moveCursorRight() {
	if w.cursorX == (256*256 - 1) {
		return
	}
	w.cursorX++
	if w.viewportX < (256*256-w.viewportWidth-1) && w.cursorX >= w.viewportX+w.viewportWidth {
		w.viewportX++
	}
}

func newWorldView(w world.World) *worldView {
	wv := new(worldView)

	wv.Box = tview.NewBox()
	wv.Box.SetBorder(true)
	wv.world = w

	//Pendiente
	reverse := tcell.StyleDefault.Reverse(true)

	wv.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		wv.Box.SetTitle(fmt.Sprintf("(%d,%d)", wv.cursorX, wv.cursorY))
		if event.Key() == tcell.KeyDown {
			wv.moveCursorDown()
		} else if event.Key() == tcell.KeyUp {
			wv.moveCursorUp()
		} else if event.Key() == tcell.KeyLeft {
			wv.moveCursorLeft()
		} else if event.Key() == tcell.KeyRight {
			wv.moveCursorRight()
		} else if event.Key() == tcell.KeyRune {
			cx := wv.cursorX / 256
			cy := wv.cursorY / 256
			ox := wv.cursorX % 256
			oy := wv.cursorY % 256
			wv.world.GetChunk(cx, cy).SetRune(ox, oy, event.Rune())
			wv.moveCursorRight()
		} else {
			return event
		}
		return nil
	})
	wv.SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		/* Update the viewport if it has changed. */
		wv.viewportWidth = width
		wv.viewportHeight = height

		/* Get the chunks that will have to be repainted. */
		topLeftChunkX, topLeftChunkY := world.GetChunkAtPos(wv.viewportX, wv.viewportY)
		bottomRightChunkX, bottomRightChunkY := world.GetChunkAtPos(wv.viewportX+width-1, wv.viewportY+height-1)

		for x := topLeftChunkX; x <= bottomRightChunkX; x++ {
			for y := topLeftChunkY; y <= bottomRightChunkY; y++ {
				wv.drawChunk(screen, x, y)
			}
		}

		screen.SetCell(wv.cursorX-wv.viewportX, wv.cursorY-wv.viewportY, reverse, ' ')
		return x, y, width, height
	})
	return wv
}

func (w *worldView) drawChunk(screen tcell.Screen, chunkX, chunkY int) {
	chunk := w.world.GetChunk(chunkX, chunkY)
	chunkTopLeftX := 256*chunkX - w.viewportX
	chunkTopLeftY := 256*chunkY - w.viewportY

	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			posX := x + chunkTopLeftX
			posY := y + chunkTopLeftY
			//TODO set a combc value
			screen.SetContent(posX+12, posY+12, chunk.GetRune(x, y), nil, tcell.StyleDefault)
		}
	}
}
