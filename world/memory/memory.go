package memory

import (
	"time"

	"github.com/dannywolfmx/iwb/world"
)

const width = 256

const height = 256

type MemoryChunk struct {
	elements   world.Elements
	lastUpdate int64
}

func NewMemoryChunk() *MemoryChunk {
	chunk := &MemoryChunk{
		elements:   make(world.Elements),
		lastUpdate: time.Now().Unix(),
	}
	return chunk
}

// LastUpdatedAt returns the timestamp at which the chunk was last updated.
func (c *MemoryChunk) LastUpdatedAt() int64 {
	return c.lastUpdate
}

// GetRune returns the rune at position (x,y)
func (c *MemoryChunk) GetElements() world.Elements {
	return c.elements
}

// GetRune returns the rune at position (x,y)
func (c *MemoryChunk) GetElement(position world.Position) world.Element {
	return c.elements[position]
}

// SetRune updates the value of a given coordinate in a chunk
func (c *MemoryChunk) SetElement(position world.Position, element world.Element) {
	c.elements[position] = element
	c.lastUpdate = time.Now().Unix()
}

// MemoryWorld stores a world in-memory. This is the basis to other
// World implementations that would use MemoryWorld as a cache until
// data is moved somewhere else (for instance a drive or HTTP server).
type MemoryWorld struct {
	chunks []*MemoryChunk
}

// NewRAMWorld creates a world storing data in RAM
func NewMemoryWorld() *MemoryWorld {
	return &MemoryWorld{
		chunks: make([]*MemoryChunk, width*height),
	}
}

func (w *MemoryWorld) GetChunk(x int, y int) world.Chunk {
	chunk := w.chunks[y*width+x]
	if chunk == nil {
		chunk = NewMemoryChunk()
		w.chunks[y*width+x] = chunk
	}
	return chunk
}

func (w *MemoryWorld) Persist() error {
	return nil
}
