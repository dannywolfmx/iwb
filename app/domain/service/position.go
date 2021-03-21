package service

import "github.com/dannywolfmx/iwb/app/domain/entity"

const maxPositionX = 100
const maxPositionY = 100
const minPositionX = 0
const minPositionY = 0

type PositionService struct{}

func (p PositionService) IsAValidPosition(position entity.Position) bool {
	x := position.X
	y := position.Y
	return x <= maxPositionX && x >= minPositionX && y <= maxPositionY && y >= minPositionY
}
