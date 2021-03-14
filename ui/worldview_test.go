package ui

import (
	"testing"

	"github.com/dannywolfmx/iwb/entity"
	"github.com/gdamore/tcell/v2"
)

type WorldTest struct {
	position entity.Position
}

func (w *WorldTest) Persist() error {
	return nil
}
func (w *WorldTest) SetPosition(viewport, chunkLocation entity.Position) {
}
func (w *WorldTest) GetPosition() (entity.Position, entity.Position) {
	return w.position, w.position
}
func (w *WorldTest) GetChunk(position entity.Position) *entity.Chunk {
	return &entity.Chunk{}
}

func TestMoveViewportX(t *testing.T) {
	testScreen := tcell.NewSimulationScreen("")
	world := &WorldTest{}
	w := NewWorldView(testScreen, world)
	positionX := 0

	//Actual position of X
	positionX = 12
	w.moveViewportX(12)

	if positionX != w.viewport.X {
		t.Fatalf("ViewportX actual porition %d, expected position %d", w.viewport.X, positionX)
	}

	positionX += 36
	w.moveViewportX(36)

	if positionX != w.viewport.X {
		t.Fatalf("ViewportX actual porition %d, expected position %d", w.viewport.X, positionX)
	}

	w.moveViewportX(-50)
	if positionX-50 != w.viewport.X {
		t.Fatalf("ViewportX actual porition %d, expected position %d", w.viewport.X, positionX)
	}
}

func TestMoveViewportY(t *testing.T) {
	testScreen := tcell.NewSimulationScreen("")
	world := &WorldTest{}
	w := NewWorldView(testScreen, world)
	positionY := 0

	//Actual position of X
	positionY = 12
	w.moveViewportY(12)

	if positionY != w.viewport.Y {
		t.Fatalf("ViewportX actual porition %d, expected position %d", w.viewport.X, positionY)
	}

	positionY += 36
	w.moveViewportY(36)

	if positionY != w.viewport.Y {
		t.Fatalf("ViewportX actual porition %d, expected position %d", w.viewport.X, positionY)
	}

	w.moveViewportY(-50)
	if positionY-50 != w.viewport.Y {
		t.Fatalf("ViewportX actual porition %d, expected position %d", w.viewport.X, positionY)
	}
}
