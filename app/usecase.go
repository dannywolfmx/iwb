//usecase.go descripbe the interface of the "app" usecases
package app

import "github.com/dannywolfmx/iwb/app/domain/entity"

type WorldUsecase interface {
	//GetChunks return a map of position as key and *Chunk as value
	GetChunk(x, y int) (*entity.Chunk, error)
	//SetChunks set a map of position as key and *chunk as value
	SetElement(chunk *entity.Chunk, x, y int, element entity.Element) error
	//GetChunk return a specific *chunk from matched with the position parameter
}
