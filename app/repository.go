package app

import "github.com/dannywolfmx/iwb/app/domain/entity"

type WorldRepository interface {
	//GetChunk return a chunk pointer of the position
	GetChunk(position entity.Position) (*entity.Chunk, error)
	SetElement(chunk *entity.Chunk, elementPosition entity.Position, element entity.Element) error
}

type CameraRepository interface {
	//Viewport
	//
	//SetViewport will set the actual position of the camera
	SetViewport(viewport entity.Position)
	//GetViewport will set the actual position of the camera
	//Return a viewport, and chunkLocation
	GetViewport() entity.Position
}

//TODO PRUEBA
type UserRepository interface {
	SetPosition(user *entity.User, position entity.Position) error
	//GetPosition find the position of the user
	//The firs return value is the user position in the chunk
	//The second return value is the chunk position in the world
	GetPosition(user *entity.User) (entity.Position, entity.Position, error)
	Exist(user *entity.User) bool
}
