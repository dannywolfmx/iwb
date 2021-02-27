package world

const width = 256

const height = 256

type Chunk interface {
	GetRune(x int, y int) rune
	SetRune(x int, y int, char rune)
	LastUpdatedAt() int64
	GetRunes() string
}

type World interface {
	GetChunk(x int, y int) Chunk
	Persist() error
}

func GetChunkAtPos(x int, y int) (int, int) {
	return x / width, y / height
}
