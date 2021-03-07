package file

import (
	"encoding/gob"
	"os"

	"github.com/dannywolfmx/iwb/world"
)

const Filename = "world.dat"

type FileWorld struct {
	Chunks              map[world.Position]*world.Chunk
	ActualViewport      world.Position
	ActualChunkLocation world.Position
}

//NewFileWorld generate a new World
func NewFileWorld() *FileWorld {
	return &FileWorld{
		Chunks: make(map[world.Position]*world.Chunk),
	}
}

func (w *FileWorld) SetPosition(viewport, chunkLocation world.Position) {
	w.ActualViewport = viewport
	w.ActualChunkLocation = chunkLocation
}

//GetPosition return the viewport, and chunk location
func (w *FileWorld) GetPosition() (world.Position, world.Position) {
	return w.ActualViewport, w.ActualChunkLocation
}

//GetChunk find a chunk in the given position or genereta a new one
func (w *FileWorld) GetChunk(position world.Position) *world.Chunk {
	//If the chunk exist return the finded chunk
	if chunk, ok := w.Chunks[position]; ok {
		return chunk
	}

	//Create a new chunk
	chunk := world.NewChunk()

	//Set the new chunk
	w.Chunks[position] = chunk

	//Return the reference
	return chunk
}

func (w *FileWorld) Persist() error {
	return SaveToFile(Filename, w)
}

//SaveToFile will save an byte slice of data into the path file
func SaveToFile(path string, world *FileWorld) error {
	file, err := os.Create(Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)

	return encoder.Encode(world)
}

//TODO check the lock variable, its the same as SaveToFIle
func LoadWorld(path string) (*FileWorld, error) {
	world := NewFileWorld()
	file, err := os.Open(Filename)
	if err != nil {
		return world, nil
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(world); err != nil {
		return nil, err
	}

	return world, nil
}
