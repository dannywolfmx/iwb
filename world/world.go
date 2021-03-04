package world

const width = 256

const height = 256

type Position struct {
	X, Y int
}

type Chunk interface {
	GetRune(position Position) rune
	SetRune(position Position, char rune)
	LastUpdatedAt() int64
}

type World interface {
	GetChunk(x int, y int) Chunk
	Persist() error
}
