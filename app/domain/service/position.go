package service

import "github.com/dannywolfmx/iwb/app/domain/entity"

const chunkSize = 256

const maxElementPositionX = 255 * chunkSize
const maxElementPositionY = maxElementPositionX
const minElementPositionX = 0
const minElementPositionY = 0

const maxUserPositionX = maxElementPositionX
const maxUserPositionY = maxElementPositionY
const minUserPositionX = minElementPositionX
const minUserPositionY = minElementPositionY

type PositionService struct{}

func (p PositionService) IsAValidChunkPosition(position entity.Position) bool {
	x := position.X
	y := position.Y
	return x < chunkSize && x >= 0 && y < chunkSize && y >= 0
}

func (p PositionService) IsAValidElementPosition(position entity.Position) bool {
	x := position.X
	y := position.Y
	return x <= maxElementPositionX && x >= minElementPositionX && y <= maxElementPositionY && y >= minElementPositionY
}

func (p PositionService) IsAValidUserPosition(position entity.Position) bool {
	x := position.X
	y := position.Y
	return x <= maxUserPositionX && x >= minUserPositionX && y <= maxUserPositionY && y >= minElementPositionY
}

func (p PositionService) GetDefaultPosition() entity.Position {
	return entity.Position{
		X: 0, Y: 0,
	}
}

//GetChunkPosition: pass the actual user position to calculate the actual chunk
//this method doen't care about validate the given user position, you need to check it before
func (p PositionService) CalculateChunkPosition(userPosition entity.Position) entity.Position {
	return entity.Position{
		X: (userPosition.X / chunkSize),
		Y: (userPosition.Y / chunkSize),
	}
}
