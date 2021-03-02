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

	if positionX != w.viewportX {
		t.Fatalf("ViewportX actual porition %d, expected position %d", w.viewportX, positionX)
	}

	positionX += 36
	w.moveViewportX(36)

	if positionX != w.viewportX {
		t.Fatalf("ViewportX actual porition %d, expected position %d", w.viewportX, positionX)
	}

	positionX += -50
	w.moveViewportX(-50)
	if positionX != w.viewportX {
		t.Fatalf("ViewportX actual porition %d, expected position %d", w.viewportX, positionX)
	}
}

func TestMoveViewportY(t *testing.T) {
	testScreen := tcell.NewSimulationScreen("")
	w := NewWorldView(testScreen)
	positionY := 0

	//Actual position of X
	positionY = 12
	w.moveViewportY(12)

	if positionY != w.viewportY {
		t.Fatalf("ViewportX actual porition %d, expected position %d", w.viewportX, positionY)
	}

	positionY += 36
	w.moveViewportY(36)

	if positionY != w.viewportY {
		t.Fatalf("ViewportX actual porition %d, expected position %d", w.viewportX, positionY)
	}

	positionY += -50
	w.moveViewportY(-50)
	if positionY != w.viewportY {
		t.Fatalf("ViewportX actual porition %d, expected position %d", w.viewportX, positionY)
	}
}
