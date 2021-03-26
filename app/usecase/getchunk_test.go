package usecase

import (
	"testing"

	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
	"github.com/dannywolfmx/iwb/app/repository/mock"
)

func TestGetChunk(t *testing.T) {
	//Set parammeters to usecase
	repo := &mock.WorldMockRepository{
		Position: entity.Position{},
		Chunk:    entity.NewChunk(),
		Element:  't',
	}

	//Validate the interface vs NewGetChunk
	var getChunk app.GetChunk

	getChunk = NewGetChunk(repo)
	//Right value
	chunk, err := getChunk.Execute(entity.Position{X: 0, Y: 0})

	if err != nil {
		t.Fatalf("Error in GetChunk %s", err)
	}

	if chunk == nil {
		t.Fatal("Chunk is nil")
	}

	//Test if the usecase is passing the position to the repositiory
	chunk, _ = getChunk.Execute(entity.Position{X: 1, Y: 0})

	//This need to be an nil value
	if chunk != nil {
		t.Fatal("Chunk is nil")
	}
}
