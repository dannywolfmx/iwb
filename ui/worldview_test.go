package ui

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestMoveViewportX(t *testing.T) {
	testScreen := tcell.NewSimulationScreen("")
	w := NewWorldView(testScreen)
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

	positionX += -50
	w.moveViewportX(-50)
	if positionX != w.viewport.X {
		t.Fatalf("ViewportX actual porition %d, expected position %d", w.viewport.X, positionX)
	}
}

func TestMoveViewportY(t *testing.T) {
	testScreen := tcell.NewSimulationScreen("")
	w := NewWorldView(testScreen)
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

	positionY += -50
	w.moveViewportY(-50)
	if positionY != w.viewport.Y {
		t.Fatalf("ViewportX actual porition %d, expected position %d", w.viewport.X, positionY)
	}
}
