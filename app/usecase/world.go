package usecase

import (
	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
	"github.com/dannywolfmx/iwb/app/domain/service"
)

type worldUsecase struct {
	repo    app.WorldRepository
	service *service.PositionService
}

//NewWorldUsecase set a repo
func NewWorldUsecase(repo app.WorldRepository) *worldUsecase {
	return &worldUsecase{
		repo:    repo,
		service: &service.PositionService{},
	}
}

//GetChunk pass a valid x, y postion and return a chunk pointer or error
func (w *worldUsecase) GetChunk(x, y int) (*entity.Chunk, error) {
	//Get Position
	position := entity.Position{X: x, Y: y}
	//Validate the position with the service
	if ok := w.service.IsAValidChunkPosition(position); !ok {
		//Invalid position
		return nil, app.ErrorInvalidPosition
	}

	//Get Chunk in the position
	chunk, err := w.repo.GetChunk(position)

	if err != nil {
		return nil, err
	}

	//Return result
	return chunk, nil
}

func (w *worldUsecase) SetElement(chunk *entity.Chunk, x, y int, element entity.Element) error {
	elementPosition := entity.Position{X: x, Y: y}
	//ValidatePosition
	//Validate the position with the service
	if ok := w.service.IsAValidElementPosition(elementPosition); !ok {
		//Invalid position
		return app.ErrorInvalidPosition
	}
	//Return error if exist or nil
	return w.repo.SetElement(chunk, elementPosition, element)
}
