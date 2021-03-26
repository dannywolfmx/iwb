package service

import (
	"testing"

	"github.com/dannywolfmx/iwb/app/domain/entity"
)

func TestIsAValidChunkPosition(t *testing.T) {
	//service
	ser := PositionService{}

	//Test max position
	validPosition := entity.Position{
		X: maxChunkPositionX,
		Y: maxChunkPositionY,
	}
	if ok := ser.IsAValidChunkPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	//Test min positon
	validPosition = entity.Position{
		X: minChunkPositionX,
		Y: minChunkPositionY,
	}
	if ok := ser.IsAValidChunkPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	validPosition = entity.Position{
		X: minChunkPositionX,
		Y: maxChunkPositionY,
	}
	if ok := ser.IsAValidChunkPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	validPosition = entity.Position{
		X: maxChunkPositionX,
		Y: minChunkPositionX,
	}
	if ok := ser.IsAValidChunkPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: maxChunkPositionX + 1,
		Y: maxChunkPositionY + 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidChunkPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: minChunkPositionX - 1,
		Y: minChunkPositionY - 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidChunkPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: minChunkPositionX,
		Y: minChunkPositionY - 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidChunkPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: minChunkPositionX - 1,
		Y: minChunkPositionY,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidChunkPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: maxChunkPositionX + 1,
		Y: minChunkPositionY,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidChunkPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: maxChunkPositionX,
		Y: maxChunkPositionY + 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidChunkPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}
}

func TestIsAValidElementPosition(t *testing.T) {
	//service
	ser := PositionService{}

	//Test max position
	validPosition := entity.Position{
		X: maxElementPositionX,
		Y: maxElementPositionY,
	}
	if ok := ser.IsAValidElementPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	//Test min positon
	validPosition = entity.Position{
		X: minElementPositionX,
		Y: minElementPositionY,
	}
	if ok := ser.IsAValidElementPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	validPosition = entity.Position{
		X: minElementPositionX,
		Y: maxElementPositionY,
	}
	if ok := ser.IsAValidElementPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	validPosition = entity.Position{
		X: maxElementPositionX,
		Y: minElementPositionX,
	}
	if ok := ser.IsAValidElementPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: maxElementPositionX + 1,
		Y: maxElementPositionY + 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidElementPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: minElementPositionX - 1,
		Y: minElementPositionY - 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidElementPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: minElementPositionX,
		Y: minElementPositionY - 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidElementPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: minElementPositionX - 1,
		Y: minElementPositionY,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidElementPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: maxElementPositionX + 1,
		Y: minElementPositionY,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidElementPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: maxElementPositionX,
		Y: maxElementPositionY + 1,
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
		X: maxUserPositionX,
		Y: maxUserPositionY,
	}
	if ok := ser.IsAValidUserPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	//Test min positon
	validPosition = entity.Position{
		X: minUserPositionX,
		Y: minUserPositionY,
	}
	if ok := ser.IsAValidUserPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	validPosition = entity.Position{
		X: minUserPositionX,
		Y: maxUserPositionY,
	}
	if ok := ser.IsAValidUserPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	validPosition = entity.Position{
		X: maxUserPositionX,
		Y: minUserPositionX,
	}
	if ok := ser.IsAValidUserPosition(validPosition); !ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: maxUserPositionX + 1,
		Y: maxUserPositionY + 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidUserPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: minUserPositionX - 1,
		Y: minUserPositionY - 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidUserPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: minUserPositionX,
		Y: minUserPositionY - 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidUserPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: minUserPositionX - 1,
		Y: minUserPositionY,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidUserPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: maxUserPositionX + 1,
		Y: minUserPositionY,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidUserPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}

	//Test invalid position
	validPosition = entity.Position{
		X: maxUserPositionX,
		Y: maxUserPositionY + 1,
	}
	//The ok value need to be a "false" value to validate the test
	if ok := ser.IsAValidUserPosition(validPosition); ok {
		t.Fatal("Should be a valid position")
	}
}
