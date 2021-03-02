package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

func DefaultStyle() tcell.Style {
	return tcell.StyleDefault
}

//newDefaultGrid will set a default data to tview.Grid struct
//
//* Show a border arround the terminal
//* Show a chess like board
func NewDefaultScreen(style tcell.Style) (tcell.Screen, error) {
	encoding.Register()
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
