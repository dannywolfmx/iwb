package file

import (
	"encoding/gob"
	"os"
	"sync"

	"github.com/dannywolfmx/iwb/world"
)

var lock sync.Mutex

type FileWorld struct {
	sync.Mutex
	chunks []*FileChunk
}

func NewFileWorld() *FileWorld {
	return &FileWorld{
		chunks: make([]*FileChunk, 1),
	}
}

func (w *FileWorld) GetChunk(x int, y int) world.Chunk {
	chunk := w.chunks[0]
	if chunk != nil {
		return chunk
	}

	// So the chunk exists and has been loaded. Decode it.

	file, err := os.Open("chunk.dat")

	if err != nil {

		chunk := NewFileChunk(make(world.Elements))
		w.chunks[0] = chunk
		return chunk
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	err = decoder.Decode(&w.chunks)

	if err != nil {
		panic(err)
	}
	return w.chunks[0]
}

func (w *FileWorld) Persist() error {
	file, err := os.Create("chunk.dat")
	if err != nil {
		return err
	}

	defer file.Close()

	encode := gob.NewEncoder(file)

	return encode.Encode(w.chunks)
}

//SaveToFile will save an byte slice of data into the path file
func SaveToFile(path string, data []byte) error {
	lock.Lock()
	defer lock.Unlock()
	return os.WriteFile(path, data, 0644)
}

//TODO check the lock variable, its the same as SaveToFIle
func LoadFromFile(path string) ([]byte, error) {
	lock.Lock()
	defer lock.Unlock()
	return os.ReadFile(path)
}
