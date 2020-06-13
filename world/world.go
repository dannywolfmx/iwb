package world

import "time"

const width = 256

const height = 256

// Chunk is a chunk
type Chunk struct {
	data       []rune
	lastUpdate int64
}

// NewChunk does things
func NewChunk() *Chunk {
	chunk := &Chunk{
		data:       make([]rune, width*height),
		lastUpdate: time.Now().Unix(),
	}
	for i := range chunk.data {
		chunk.data[i] = ' '
	}
	return chunk
}

// LastUpdatedAt returns the timestamp at which the chunk was last updated.
func (c *Chunk) LastUpdatedAt() int64 {
	return c.lastUpdate
}

// GetRunes gets all the data for this chunk
func (c *Chunk) GetRunes() string {
	return string(c.data)
}

// GetRune returns the rune at position (x,y)
func (c *Chunk) GetRune(x int32, y int32) rune {
	return c.data[y*width+x]
}

// SetRune updates the value of a given coordinate in a chunk
func (c *Chunk) SetRune(x int32, y int32, char rune) {
	c.data[y*width+x] = char
	c.lastUpdate = time.Now().Unix()
}

// World is world
type World struct {
	chunks []*Chunk
}

func NewWorld() *World {
	return &World{
		chunks: make([]*Chunk, width*height),
	}
}

func (w *World) GetChunk(x int32, y int32) *Chunk {
	chunk := w.chunks[y*width+x]
	if chunk == nil {
		chunk = NewChunk()
		w.chunks[y*width+x] = chunk
	}
	return chunk
}

func GetChunkAtPos(x int, y int) (int, int) {
	return x / width, y / height
}
