package stream

import (
	"encoding/gob"
	"io"
	"os"

	"github.com/dannywolfmx/iwb/app/domain/entity"
)

//world is the struct repository
type world struct {
	cacheChunks map[entity.Position]*entity.Chunk
}

//NewWorldRepository will try to read the stream to load
//the chunks cache and return it with the world.
//if the stream can't get a proper data, the world will be generate with
//a empty cache and save it in the stream
func NewWorldRepository(stream io.ReadWriter) *world {
	//Generate a new chunk
	chunk := make(map[entity.Position]*entity.Chunk)
	if err := load(stream, &chunk); err != nil {
		//Generate a new chunk becouse the other are overwritten
		chunk = make(map[entity.Position]*entity.Chunk)
	}

	//Return a new
	return &world{
		cacheChunks: chunk,
	}
}

func (w *world) GetChunk(position entity.Position) *entity.Chunk {
	//Check if the chunk exist in the cache
	if chunk, ok := w.cacheChunks[position]; ok {
		return chunk
	}

	//Return a new empty chunk
	return entity.NewChunk()
}

const Filename = "../world.dat"

type FileWorld struct {
	Chunks              map[world.Position]*world.Chunk
	ActualViewport      world.Position
	ActualChunkLocation world.Position
}

//NewFileWorld generate a new World
func NewFileWorld() *FileWorld {
	chunks := make(map[world.Position]*world.Chunk)
	return &FileWorld{
		Chunks: chunks,
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
