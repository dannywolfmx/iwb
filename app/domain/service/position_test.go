package service

import (
	"testing"

	"github.com/dannywolfmx/iwb/app/domain/entity"
)

func TestIsAValidPosition(t *testing.T) {
	//service
	ser := PositionService{}

	//Test max position
	validPosition := entity.Position{
		X: maxPositionX,
		Y: maxPositionY,
	}
	if ok := ser.IsAValidPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	//Test min positon
	validPosition = entity.Position{
		X: minPositionX,
		Y: minPositionY,
	}
	if ok := ser.IsAValidPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	validPosition = entity.Position{
		X: minPositionX,
		Y: maxPositionY,
	}
	if ok := ser.IsAValidPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	validPosition = entity.Position{
		X: maxPositionX,
		Y: minPositionX,
	}
	if ok := ser.IsAValidPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: maxPositionX + 1,
		Y: maxPositionY + 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: minPositionX - 1,
		Y: minPositionY - 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: minPositionX,
		Y: minPositionY - 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: minPositionX - 1,
		Y: minPositionY,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: maxPositionX + 1,
		Y: minPositionY,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: maxPositionX,
		Y: maxPositionY + 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}
}
