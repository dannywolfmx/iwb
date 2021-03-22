//usecase.go descripbe the interface of the "app" usecases
package app

import "github.com/dannywolfmx/iwb/app/domain/entity"

//SetChunks set a map of position as key and *chunk as value
type SetElement interface {
	//Execute the usecase
	Execute(chunk *entity.Chunk, x, y int, element entity.Element) error
}

//GetChunks return a map of position as key and *Chunk as value
type GetChunk interface {
	//Execute the usecase
	Execute(x, y int) (*entity.Chunk, error)
}
