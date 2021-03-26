//usecase.go descripbe the interface of the "app" usecases
package app

import "github.com/dannywolfmx/iwb/app/domain/entity"

//World interfaces

//SetChunks set a map of position as key and *chunk as value
type SetElement interface {
	//Execute the usecase
	Execute(chunk *entity.Chunk, x, y int, element entity.Element) error
}

//GetChunks return a map of position as key and *Chunk as value
type GetChunk interface {
	//Execute the usecase
	Execute(position entity.Position) (*entity.Chunk, error)
}

//
//User (Session) interfaces
//

//SetUserPosition
//Set a position of a current user into the actual world chunk
type SetUserPosition interface {
	//Execute the usecase
	Execute(user *entity.User, position entity.Position) error
}

// TODO Implement
//
//Get the current user posicion into the chunk and chunk position into de world
//return
//	1 - user position into the chunk
//	2 - chunk position into the world
//	3 - error
//
type GetUserPosition interface {
	//Execute the usecase
	Execute(user *entity.User) (entity.Position, entity.Position, error)
}
