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

func NewWorldUsecase(repo app.WorldRepository, service *service.PositionService) *worldUsecase {
	return &worldUsecase{
		repo:    repo,
		service: service,
	}
}

func (w *worldUsecase) GetChunk(x, y int) (*entity.Chunk, error) {
	//Get Position
	position := entity.Position{
		X: x, Y: y,
	}
	//Validate the position with the service
	if ok := w.service.IsAValidPosition(position); !ok {
		//Invalid position
		return nil, app.InvalidPosition
	}

	//Get Chunk in the position
	chunk := w.repo.GetChunk(position)
	//Enviar chunk al cliente
	return chunk, nil
}
