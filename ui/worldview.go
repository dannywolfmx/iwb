package ui

import (
	"fmt"

	"danirod.es/pkg/iwb/world"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type WorldView struct {
	*tview.Box
	cursorX, cursorY              int
	viewportX, viewportY          int
	viewportWidth, viewportHeight int
	world                         *world.World
}

func (w *WorldView) moveCursorUp() {
	if w.cursorY == 0 {
		return
	}
	w.cursorY--
	if w.viewportY > 0 && w.cursorY < w.viewportY {
		w.viewportY--
	}
}

func (w *WorldView) moveCursorDown() {
	if w.cursorY == (256*256 - 1) {
		return
	}
	w.cursorY++
	if w.viewportY < (256*256-w.viewportHeight-1) && w.cursorY >= w.viewportY+w.viewportHeight {
		w.viewportY++
	}
}

func (w *WorldView) moveCursorLeft() {
	if w.cursorX == 0 {
		return
	}
	w.cursorX--
	if w.viewportX >= 0 && w.cursorX < w.viewportX {
		w.viewportX--
	}
}

func (w *WorldView) moveCursorRight() {
	if w.cursorX == (256*256 - 1) {
		return
	}
	w.cursorX++
	if w.viewportX < (256*256-w.viewportWidth-1) && w.cursorX >= w.viewportX+w.viewportWidth {
		w.viewportX++
	}
}

func NewWorldView(w *world.World) *WorldView {
	wv := new(WorldView)
	wv.Box = tview.NewBox()
	wv.Box.SetBorder(true)
	wv.world = w
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
			cx := int32(wv.cursorX / 256)
			cy := int32(wv.cursorY / 256)
			ox := int32(wv.cursorX % 256)
			oy := int32(wv.cursorY % 256)
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

func (w *WorldView) drawChunk(screen tcell.Screen, chunkX, chunkY int) {
	chunk := w.world.GetChunk(int32(chunkX), int32(chunkY))
	chunkTopLeftX := 256*chunkX - w.viewportX
	chunkTopLeftY := 256*chunkY - w.viewportY

	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			posX := x + chunkTopLeftX
			posY := y + chunkTopLeftY
			screen.SetCell(posX, posY, tcell.StyleDefault, chunk.GetRune(int32(x), int32(y)))
		}
	}
}
