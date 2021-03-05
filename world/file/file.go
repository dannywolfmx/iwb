package file

import (
	"github.com/dannywolfmx/iwb/world"
	"github.com/dannywolfmx/iwb/world/memory"
)

type FileChunk struct {
	*memory.MemoryChunk
}

//TODO Pass change element to *element
func NewFileChunk(elements world.Elements) *FileChunk {
	return &FileChunk{
		MemoryChunk: memory.NewMemoryChunk(elements),
	}
}
