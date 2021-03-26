package service

import "github.com/dannywolfmx/iwb/app/domain/entity"

const maxChunkPositionX = 100
const maxChunkPositionY = 100
const minChunkPositionX = 0
const minChunkPositionY = 0

const maxElementPositionX = 100
const maxElementPositionY = 100
const minElementPositionX = 0
const minElementPositionY = 0

const maxUserPositionX = 100
const maxUserPositionY = 100
const minUserPositionX = 0
const minUserPositionY = 0

type PositionService struct{}

func (p PositionService) IsAValidChunkPosition(position entity.Position) bool {
	x := position.X
	y := position.Y
	return x <= maxChunkPositionX && x >= minChunkPositionX && y <= maxChunkPositionY && y >= minChunkPositionY
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
