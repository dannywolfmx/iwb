package usecase

import (
	"os"
	"testing"

	"github.com/dannywolfmx/iwb/app/domain/entity"
	"github.com/dannywolfmx/iwb/app/domain/service"
)

type TestRepository struct {
	position entity.Position
}

func (t *TestRepository) GetChunk(position entity.Position) *entity.Chunk {
	//Check the position to be the same
	if t.position == position {
		return entity.NewChunk()
	}
	//Return nil as chunk to represent an error in the validate
	return nil
}

var usecase *worldUsecase

func TestMain(m *testing.M) {
	//Setup
	repo := &TestRepository{}
	service := &service.PositionService{}
	usecase = NewWorldUsecase(repo, service)
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
	chunk, err = usecase.GetChunk(1, 0)

	if err != nil {
		t.Fatalf("Error in GetChunk %s", err)
	}

	//This need to be an nil value
	if chunk != nil {
		t.Fatal("Chunk is nil")
	}
}
