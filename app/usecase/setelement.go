package usecase

import (
	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
	"github.com/dannywolfmx/iwb/app/domain/service"
)

type setElement struct {
	repo    app.WorldRepository
	service *service.PositionService
}

//NewWorldUsecase set a repo
func NewSetElement(repo app.WorldRepository) *setElement {
	return &setElement{
		repo:    repo,
		service: &service.PositionService{},
	}
}

func (w *setElement) Execute(chunk *entity.Chunk, x, y int, element entity.Element) error {
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
