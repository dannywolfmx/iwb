package file

import (
	"encoding/gob"
	"os"

	"github.com/dannywolfmx/iwb/world"
)

const Filename = "world.dat"

type FileWorld struct {
	Chunks         []*world.Chunk
	ActualPosition world.Position
}

//NewFileWorld generate a new World
func NewFileWorld() *FileWorld {
	return &FileWorld{
		Chunks: []*world.Chunk{
			world.NewChunk(),
		},
	}
}

func (w *FileWorld) SetPosition(position world.Position) {
	w.ActualPosition = position
}

func (w *FileWorld) GetPosition() world.Position {
	return w.ActualPosition
}

func (w *FileWorld) GetChunk(x, y int) *world.Chunk {
	return w.Chunks[0]
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
