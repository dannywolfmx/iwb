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
