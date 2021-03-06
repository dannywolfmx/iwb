package memory

import (
	"github.com/dannywolfmx/iwb/world"
)

const width = 256
const height = 256

// MemoryWorld stores a world in-memory. This is the basis to other
// World implementations that would use MemoryWorld as a cache until
// data is moved somewhere else (for instance a drive or HTTP server).
type MemoryWorld struct {
	chunks []*world.Chunk
}

// NewRAMWorld creates a world storing data in RAM
func NewMemoryWorld() *MemoryWorld {
	return &MemoryWorld{
		chunks: make([]*world.Chunk, width*height),
	}
}

func (w *MemoryWorld) GetChunk(x int, y int) *world.Chunk {
	chunk := w.chunks[y*width+x]
	if chunk == nil {
		chunk = world.NewChunk()
		w.chunks[y*width+x] = chunk
	}
	return chunk
}
