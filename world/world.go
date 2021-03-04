package world

const width = 256

const height = 256

type Position struct {
	X, Y int
}

type Element = rune
type Elements = map[Position]rune

type Chunk interface {
	GetElement(position Position) Element
	GetElements() Elements
	SetElement(position Position, element Element)
	LastUpdatedAt() int64
}

type World interface {
	GetChunk(x int, y int) Chunk
	Persist() error
}
