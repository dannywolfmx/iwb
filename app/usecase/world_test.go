package usecase

import (
	"os"
	"testing"

	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
)

type TestRepository struct {
	position entity.Position
	chunk    *entity.Chunk
	element  entity.Element
}

func (t *TestRepository) GetChunk(position entity.Position) (*entity.Chunk, error) {
	//Check the position to be the same
	if t.position == position {
		return entity.NewChunk(), nil
	}
	//Return nil as chunk to represent an error in the validate
	return nil, app.ErrorInvalidPosition
}

func (t *TestRepository) SetElement(chunk *entity.Chunk, position entity.Position, element entity.Element) error {
	if t.position != position || t.element != element {
		return app.ErrorOnSetElementToDB
	}
	return nil
}

var usecase *worldUsecase

func TestMain(m *testing.M) {
	//Setup
	repo := &TestRepository{
		position: entity.Position{},
		chunk:    entity.NewChunk(),
		element:  't',
	}

	usecase = NewWorldUsecase(repo)
	//run
	status := m.Run()
	//
	//Clean
	//
	os.Exit(status)
}

func TestGetChunk(t *testing.T) {
	//Right value
	chunk, err := usecase.GetChunk(0, 0)

	if err != nil {
		t.Fatalf("Error in GetChunk %s", err)
	}

	if chunk == nil {
		t.Fatal("Chunk is nil")
	}

	//Test if the usecase is passing the position to the repositiory
	chunk, _ = usecase.GetChunk(1, 0)

	//This need to be an nil value
	if chunk != nil {
		t.Fatal("Chunk is nil")
	}
}

func TestSetElement(t *testing.T) {
	//Right value
	err := usecase.SetElement(&entity.Chunk{}, 0, 0, 't')

	if err != nil {
		t.Fatalf("Error in SetChunk %s", err)
	}

	//Test if the usecase is passing the position to the repositiory
	err = usecase.SetElement(&entity.Chunk{}, 0, 0, 'r')

	//This test need to fail to be right
	if err == nil {
		t.Fatal("Test need to get an error")
	}

	//Test if the usecase is passing the position to the repositiory
	err = usecase.SetElement(&entity.Chunk{}, 1, 0, 't')

	//This test need to fail to be right
	if err == nil {
		t.Fatal("Test need to get an error")
	}
}
