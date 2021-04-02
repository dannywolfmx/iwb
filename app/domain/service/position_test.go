package service

import (
	"testing"

	"github.com/dannywolfmx/iwb/app/domain/entity"
)

func TestIsAValidElementPosition(t *testing.T) {
	//service
	ser := PositionService{}

	//Test max position
	validPosition := entity.Position{
		X: MaxElementPositionX,
		Y: MaxElementPositionY,
	}
	if ok := ser.IsAValidElementPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	//Test min positon
	validPosition = entity.Position{
		X: MinElementPositionX,
		Y: MinElementPositionY,
	}
	if ok := ser.IsAValidElementPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	validPosition = entity.Position{
		X: MinElementPositionX,
		Y: MaxElementPositionY,
	}
	if ok := ser.IsAValidElementPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	validPosition = entity.Position{
		X: MaxElementPositionX,
		Y: MinElementPositionX,
	}
	if ok := ser.IsAValidElementPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: MaxElementPositionX + 1,
		Y: MaxElementPositionY + 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidElementPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: MinElementPositionX - 1,
		Y: MinElementPositionY - 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidElementPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: MinElementPositionX,
		Y: MinElementPositionY - 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidElementPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: MinElementPositionX - 1,
		Y: MinElementPositionY,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidElementPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: MaxElementPositionX + 1,
		Y: MinElementPositionY,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidElementPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: MaxElementPositionX,
		Y: MaxElementPositionY + 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidElementPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}
}

func TestIsAValidUserPosition(t *testing.T) {
	//service
	ser := PositionService{}

	//Test max position
	validPosition := entity.Position{
		X: MaxUserPositionX,
		Y: MaxUserPositionY,
	}
	if ok := ser.IsAValidUserPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	//Test min positon
	validPosition = entity.Position{
		X: MinUserPositionX,
		Y: MinUserPositionY,
	}
	if ok := ser.IsAValidUserPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	validPosition = entity.Position{
		X: MinUserPositionX,
		Y: MaxUserPositionY,
	}
	if ok := ser.IsAValidUserPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	validPosition = entity.Position{
		X: MaxUserPositionX,
		Y: MinUserPositionX,
	}
	if ok := ser.IsAValidUserPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: MaxUserPositionX + 1,
		Y: MaxUserPositionY + 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidUserPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: MinUserPositionX - 1,
		Y: MinUserPositionY - 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidUserPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: MinUserPositionX,
		Y: MinUserPositionY - 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidUserPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: MinUserPositionX - 1,
		Y: MinUserPositionY,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidUserPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: MaxUserPositionX + 1,
		Y: MinUserPositionY,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidUserPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: MaxUserPositionX,
		Y: MaxUserPositionY + 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidUserPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}
}

func TestGetChunkPosition(t *testing.T) {
	testUserPosition := entity.Position{
		X: MaxElementPositionX,
		Y: MaxElementPositionY,
	}

	expectedChunkPosition := entity.Position{
		X: 255,
		Y: 255,
	}

	service := PositionService{}

	chunkPosition := service.CalculateChunkPosition(testUserPosition)

	if chunkPosition != expectedChunkPosition {
		t.Fatalf("ChunkPosition %v diferent to expected position %v",
			chunkPosition, expectedChunkPosition)
	}

	testUserPosition = entity.Position{
		X: MinElementPositionX,
		Y: MinElementPositionY,
	}

	expectedChunkPosition = entity.Position{
		X: 0,
		Y: 0,
	}

	chunkPosition = service.CalculateChunkPosition(testUserPosition)

	if chunkPosition != expectedChunkPosition {
		t.Fatalf("ChunkPosition %v diferent to expected position %v",
			chunkPosition, expectedChunkPosition)
	}
}
