package world

const width = 256

const height = 256

type Chunk interface {
	GetRune(x int32, y int32) rune
	SetRune(x int32, y int32, char rune)
	LastUpdatedAt() int64
	GetRunes() string
}

type World interface {
	GetChunk(x int32, y int32) Chunk
}

func GetChunkAtPos(x int, y int) (int, int) {
	return x / width, y / height
}
