package file

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"

	"danirod.es/pkg/iwb/world"
	"danirod.es/pkg/iwb/world/memory"
)

const width = 256

const height = 256

type FileChunk struct {
	*memory.MemoryChunk
	root *FileWorld
	id   int
}

func NewFileChunk(w *FileWorld, id int) *FileChunk {
	return &FileChunk{
		MemoryChunk: memory.NewMemoryChunk(),
		root:        w,
		id:          id,
	}
}

func (f *FileChunk) SetRune(x int32, y int32, char rune) {
	f.MemoryChunk.SetRune(x, y, char)
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
	runes := make([]rune, width*height)
	buffer := bytes.NewReader(data)
	err := binary.Read(buffer, binary.LittleEndian, &runes)
	if err != nil {
		return nil, err
	}

	chunk := NewFileChunk(w, id)
	chunk.SetData(runes)
	return chunk, nil
}

func encodeBytes(chunk *FileChunk) ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.LittleEndian, chunk.GetData())
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
		chunk := NewFileChunk(w, y*width+x)
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
	for k, _ := range w.dirtyChunks {
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
