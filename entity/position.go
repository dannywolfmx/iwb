package entity

//Position will contain the dictior of a point
type Position struct {
	X, Y int
}

//Return a new position
func NewPosition(x, y int) Position {
	return Position{x, y}
}

//Set a valid position
func (p Position) SetPosition(x, y int) {
	p.X, p.Y = x, y
}

//Get a position
func (p Position) GetPosition() (int, int) {
	return p.X, p.Y
}
