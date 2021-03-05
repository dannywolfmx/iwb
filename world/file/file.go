package file

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"

	"github.com/dannywolfmx/iwb/world"
	"github.com/dannywolfmx/iwb/world/memory"
)

const width = 256

const height = 256

type FileChunk struct {
	*memory.MemoryChunk
	root *FileWorld
	id   int
}

//TODO Pass change element to *element
func NewFileChunk(w *FileWorld, id int, elements world.Elements) *FileChunk {
	return &FileChunk{
		MemoryChunk: memory.NewMemoryChunk(elements),
		root:        w,
		id:          id,
	}
}

func (f *FileChunk) SetElement(position world.Position, element world.Element) {
	f.MemoryChunk.SetElement(position, element)
	f.root.dirtyChunks[f.id] = true
}

type FileWorld struct {
	chunks      []*FileChunk
	dirtyChunks map[int]bool
}

func NewFileWorld() *FileWorld {
	return &FileWorld{
		chunks:      make([]*FileChunk, width*height),
		dirtyChunks: map[int]bool{},
	}
}

func decodeBytes(data []byte, w *FileWorld, id int) (*FileChunk, error) {
	elements := make(world.Elements)
	buffer := bytes.NewReader(data)
	err := binary.Read(buffer, binary.LittleEndian, &elements)
	if err != nil {
		return nil, err
	}

	chunk := NewFileChunk(w, id, elements)
	return chunk, nil
}

func encodeBytes(chunk *FileChunk) ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.LittleEndian, chunk.GetElements())
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (w *FileWorld) GetChunk(x int, y int) world.Chunk {
	chunk := w.chunks[y*width+x]
	if chunk != nil {
		return chunk
	}

	// So the chunk is not in memory, load from disk.
	fileName := fmt.Sprintf("%d.dat", y*width+x)
	content, err := ioutil.ReadFile(fileName)

	// So the chunk either does not exist or file is corrupt. Recreate.
	if err != nil {

		chunk := NewFileChunk(w, y*width+x, make(world.Elements))
		w.chunks[y*width+x] = chunk
		return chunk
	}

	// So the chunk exists and has been loaded. Decode it.
	chunk, err = decodeBytes(content, w, y*width+x)
	if err != nil {
		panic(err)
	}
	w.chunks[y*width+x] = chunk
	return chunk
}

func (w *FileWorld) Persist() error {
	for k := range w.dirtyChunks {
		chunk := w.chunks[k]
		bytes, err := encodeBytes(chunk)
		if err != nil {
			return err
		}

		fileName := fmt.Sprintf("%d.dat", k)
		err = ioutil.WriteFile(fileName, bytes, 0644)
		if err != nil {
			return err
		}
	}

	w.dirtyChunks = map[int]bool{}
	return nil
}
