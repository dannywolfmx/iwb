package usecase

import (
	"testing"

	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
	"github.com/dannywolfmx/iwb/app/repository/mock"
)

func TestSetElement(t *testing.T) {
	repo := &mock.WorldMockRepository{
		Position: entity.Position{},
		Chunk:    entity.NewChunk(),
		Element:  't',
	}

	//Validate the interface vs NewGetChunk
	var setElement app.SetElement
	setElement = NewSetElement(repo)
	//Right value
	err := setElement.Execute(&entity.Chunk{}, 0, 0, 't')

	if err != nil {
		t.Fatalf("Error in SetChunk %s", err)
	}

	//Test if the usecase is passing the position to the repositiory
	err = setElement.Execute(&entity.Chunk{}, 0, 0, 'r')

	//This test need to fail to be right
	if err == nil {
		t.Fatal("Test need to get an error")
	}

	//Test if the usecase is passing the position to the repositiory
	err = setElement.Execute(&entity.Chunk{}, 1, 0, 't')

	//This test need to fail to be right
	if err == nil {
		t.Fatal("Test need to get an error")
	}
}
