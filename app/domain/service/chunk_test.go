package service

import (
	"testing"

	"github.com/dannywolfmx/iwb/app/domain/entity"
)

func TestAddElementToUserChunk(t *testing.T) {

	testChunk := entity.NewChunk()
	service := ChunkService{}
	testID := "1"
	testElement := 'i'
	testPosition := entity.Position{X: 0, Y: 0}

	service.AddElementToUserChunk(testID, testElement, testPosition, testChunk)

	if len(testChunk.UsersElements) == 0 {
		t.Fatalf("Error the chunk doesn't set a new element")
	}
}
