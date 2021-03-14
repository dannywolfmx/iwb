package app

import "github.com/dannywolfmx/iwb/app/domain/entity"

type World interface {
	SetPosition()
}

type Camera interface {
	//Viewport
	//
	//SetViewport will set the actual position of the camera
	SetViewport(viewport entity.Position)
	//GetViewport will set the actual position of the camera
	//Return a viewport, and chunkLocation
	GetViewport() entity.Position
}
