//usecase.go descripbe the interface of the "app" usecases
package app

type WorldUsecase interface {
	//GetChunks return a map of position as key and *Chunk as value

	//SetChunks set a map of position as key and *chunk as value

	//GetChunk return a specific *chunk from matched with the position parameter
}
