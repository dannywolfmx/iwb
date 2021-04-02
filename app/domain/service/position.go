package service

import "github.com/dannywolfmx/iwb/app/domain/entity"

const ChunkSize = 256

const MaxElementPositionX = 255 * ChunkSize
const MaxElementPositionY = MaxElementPositionX
const MinElementPositionX = 0
const MinElementPositionY = 0

const MaxUserPositionX = MaxElementPositionX
const MaxUserPositionY = MaxElementPositionY
const MinUserPositionX = MinElementPositionX
const MinUserPositionY = MinElementPositionY

type PositionService struct{}

func (p PositionService) IsAValidChunkPosition(position entity.Position) bool {
	x := position.X
	y := position.Y
	return x < ChunkSize && x >= 0 && y < ChunkSize && y >= 0
}

func (p PositionService) IsAValidElementPosition(position entity.Position) bool {
	x := position.X
	y := position.Y
	return x <= MaxElementPositionX && x >= MinElementPositionX && y <= MaxElementPositionY && y >= MinElementPositionY
}

func (p PositionService) IsAValidUserPosition(position entity.Position) bool {
	x := position.X
	y := position.Y
	return x <= MaxUserPositionX && x >= MinUserPositionX && y <= MaxUserPositionY && y >= MinElementPositionY
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
		X: (userPosition.X / ChunkSize),
		Y: (userPosition.Y / ChunkSize),
	}
}
