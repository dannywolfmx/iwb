package world

import "time"

// ramChunk is a chunk
type ramChunk struct {
	data       []rune
	lastUpdate int64
}

func newRAMChunk() *ramChunk {
	chunk := &ramChunk{
		data:       make([]rune, width*height),
		lastUpdate: time.Now().Unix(),
	}
	for i := range chunk.data {
		chunk.data[i] = ' '
	}
	return chunk
}

// LastUpdatedAt returns the timestamp at which the chunk was last updated.
func (c *ramChunk) LastUpdatedAt() int64 {
	return c.lastUpdate
}

// GetRunes gets all the data for this chunk
func (c *ramChunk) GetRunes() string {
	return string(c.data)
}

// GetRune returns the rune at position (x,y)
func (c *ramChunk) GetRune(x int32, y int32) rune {
	return c.data[y*width+x]
}

// SetRune updates the value of a given coordinate in a chunk
func (c *ramChunk) SetRune(x int32, y int32, char rune) {
	c.data[y*width+x] = char
	c.lastUpdate = time.Now().Unix()
}

// ramWorld is world
type ramWorld struct {
	chunks []*ramChunk
}

// NewRAMWorld creates a world storing data in RAM
func NewRAMWorld() World {
	return &ramWorld{
		chunks: make([]*ramChunk, width*height),
	}
}

func (w *ramWorld) GetChunk(x int32, y int32) Chunk {
	chunk := w.chunks[y*width+x]
	if chunk == nil {
		chunk = newRAMChunk()
		w.chunks[y*width+x] = chunk
	}
	return chunk
}
