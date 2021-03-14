package ui

import (
	"fmt"
	"os"

	"github.com/dannywolfmx/iwb/entity"
	"github.com/gdamore/tcell/v2"
)

type worldView struct {
	screen        tcell.Screen
	viewport      entity.Position
	world         entity.PersistantWorld
	actualChunk   *entity.Chunk
	chunkPosition entity.Position
}

//NewWorldView create a worldView
func NewWorldView(screen tcell.Screen, w entity.PersistantWorld) *worldView {
	viewport, chunkPosition := w.GetPosition()
	return &worldView{
		screen:        screen,
		viewport:      viewport,
		world:         w,
		actualChunk:   w.GetChunk(chunkPosition),
		chunkPosition: chunkPosition,
	}
}

//Clear will erase any character into the world screen
func (w *worldView) Clear() {
	w.screen.Clear()
}

func (w *worldView) moveViewportX(position int) {
	//If a use the position like uint i will lose the negative numbers, thats because i need to do this
	viewport := int(w.viewport.X) + position
	if viewport > 255 {
		w.chunkPosition.X++
		w.actualChunk = w.world.GetChunk(w.chunkPosition)
	} else if viewport < 0 {
		w.chunkPosition.X--
		w.actualChunk = w.world.GetChunk(w.chunkPosition)
	}
	w.viewport.X = viewport
}

func (w *worldView) moveViewportY(position int) {
	//If a use the position like uint i will lose the negative numbers, thats because i need to do this
	viewport := int(w.viewport.Y) + position
	if viewport > 255 {
		w.chunkPosition.Y++
		w.actualChunk = w.world.GetChunk(w.chunkPosition)
	} else if viewport < 0 {
		w.chunkPosition.Y--
		w.actualChunk = w.world.GetChunk(w.chunkPosition)
	}
	w.viewport.Y = viewport
}

//TODO the printer dont works with special characters, just support 1 rune at the time
//TODO Deal with the uint position: if i did 1 - 2 will be 255
func (w *worldView) printOnScreen(text rune, comb []rune, viewport entity.Position, wv, hv int, style tcell.Style, moveX, moveY int) {
	positionX := int(viewport.X) - int(w.viewport.X) + moveX
	positionY := int(viewport.Y) - int(w.viewport.Y) + moveY
	//Print On Center of screen

	w.screen.SetContent(positionX+wv/2, positionY+hv/2, text, nil, style)
	//Move the position to the next rune
}

func (w *worldView) Draw() {
	w.screen.Clear()

	wv, hv := w.screen.Size()

	printBorders(w, wv, hv)
	printChunks(w, wv, hv)

	for viewport, text := range w.actualChunk.GetElements() {
		w.printOnScreen(text, nil, viewport, wv, hv, tcell.StyleDefault.Normal(), 0, 0)
	}
	//Print cursor
	w.printOnScreen(' ', nil, w.viewport, wv, hv, tcell.StyleDefault.Normal().Reverse(true), 0, 0)

	//Print title
	for i, text := range fmt.Sprintf("Chunk (%d,%d) Viewport (%d, %d)", w.chunkPosition.X, w.chunkPosition.Y, w.viewport.X, w.viewport.Y) {
		w.screen.SetContent(i+wv/2, 0, text, nil, tcell.StyleDefault)
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
				w.world.SetPosition(w.viewport, w.chunkPosition)
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

//generateBorder get a border of the actual chunk
//TODO: create a better border system using the actual viewport to eliminate unnecessary runes
func generateBorder() entity.Elements {
	border := make(entity.Elements)

	//TOP
	for i := 0; i <= 255; i++ {
		position := entity.Position{X: i, Y: 0}
		border[position] = ' '
	}

	//BOTTOM
	for i := 0; i <= 255; i++ {
		position := entity.Position{X: i, Y: 255}
		border[position] = ' '
	}

	//LEFT
	for i := 1; i <= 254; i++ {
		position := entity.Position{X: 0, Y: i}
		border[position] = ' '
	}

	//RIGHT
	for i := 1; i <= 254; i++ {
		position := entity.Position{X: 255, Y: i}
		border[position] = ' '
	}

	return border
}

func printBorders(w *worldView, wv, hv int) {
	viewDistance := 4
	for viewport, text := range generateBorder() {
		//CENTRAL BORDER (MAIN)
		//Always show it
		w.printOnScreen(text, nil, viewport, wv, hv, tcell.StyleDefault.Reverse(true), 0, 0)

		if w.viewport.Y < viewDistance {
			//TOP Border
			w.printOnScreen(text, nil, viewport, wv, hv, tcell.StyleDefault.Reverse(true), 0, -256)

			//LEFT TOP Border
			if w.viewport.X < viewDistance {
				w.printOnScreen(text, nil, viewport, wv, hv, tcell.StyleDefault.Reverse(true), -256, -256)
			} else if w.viewport.X > 255-viewDistance {
				//RIGHT TOP BORDER
				w.printOnScreen(text, nil, viewport, wv, hv, tcell.StyleDefault.Reverse(true), 256, -256)
			}
		}

		if w.viewport.Y > 255-viewDistance {
			//BOTTOM Border
			w.printOnScreen(text, nil, viewport, wv, hv, tcell.StyleDefault.Reverse(true), 0, 256)
			if w.viewport.X < viewDistance {
				//LEFT BOTTOM Border
				w.printOnScreen(text, nil, viewport, wv, hv, tcell.StyleDefault.Reverse(true), -256, 256)
			} else if w.viewport.X > 255-viewDistance {
				//RIGHT BOTTOM BORDER
				w.printOnScreen(text, nil, viewport, wv, hv, tcell.StyleDefault.Reverse(true), 256, 256)
			}
		}

		if w.viewport.X > 255-viewDistance {
			//RIGHT BORDER
			w.printOnScreen(text, nil, viewport, wv, hv, tcell.StyleDefault.Reverse(true), 256, 0)
		}

		if w.viewport.X < viewDistance {
			//LEFT Border
			w.printOnScreen(text, nil, viewport, wv, hv, tcell.StyleDefault.Reverse(true), -256, 0)
		}
	}
}

func printChunks(w *worldView, wv, hv int) {
	viewDistance := 4
	//Main chunk
	for viewport, text := range w.actualChunk.GetElements() {
		w.printOnScreen(text, nil, viewport, wv, hv, tcell.StyleDefault.Normal(), 0, 0)
	}

	if w.viewport.X < viewDistance {

		//Left chunk
		position := entity.Position{
			X: w.chunkPosition.X - 1,
			Y: w.chunkPosition.Y,
		}

		chunk := w.world.GetChunk(position)
		for viewport, text := range chunk.GetElements() {
			w.printOnScreen(text, nil, viewport, wv, hv, tcell.StyleDefault.Normal(), -256, 0)
		}
	}

	if w.viewport.X > 255-viewDistance {
		//Right chunk
		position := entity.Position{
			X: w.chunkPosition.X + 1,
			Y: w.chunkPosition.Y,
		}
		chunk := w.world.GetChunk(position)
		for viewport, text := range chunk.GetElements() {
			w.printOnScreen(text, nil, viewport, wv, hv, tcell.StyleDefault.Normal(), 256, 0)
		}
	}

	if w.viewport.Y < viewDistance {
		//TOP chunk
		position := entity.Position{
			X: w.chunkPosition.X,
			Y: w.chunkPosition.Y - 1,
		}

		chunk := w.world.GetChunk(position)
		for viewport, text := range chunk.GetElements() {
			w.printOnScreen(text, nil, viewport, wv, hv, tcell.StyleDefault.Normal(), 0, -256)
		}
	}

	if w.viewport.Y > 255-viewDistance {
		//BOTTOM chunk
		position := entity.Position{
			X: w.chunkPosition.X,
			Y: w.chunkPosition.Y + 1,
		}
		chunk := w.world.GetChunk(position)
		for viewport, text := range chunk.GetElements() {
			w.printOnScreen(text, nil, viewport, wv, hv, tcell.StyleDefault.Normal(), 0, 256)
		}
	}
}
