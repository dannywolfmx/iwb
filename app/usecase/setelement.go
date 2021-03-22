package usecase

import (
	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
	"github.com/dannywolfmx/iwb/app/domain/service"
)

//SetChunks set a map of position as key and *chunk as value
type setElement struct {
	repo    app.WorldRepository
	service *service.PositionService
}

func NewSetElement(repo app.WorldRepository) *setElement {
	return &setElement{
		repo:    repo,
		service: &service.PositionService{},
	}
}

//Execute the usecase
func (s *setElement) Execute(chunk *entity.Chunk, x, y int, element entity.Element) error {
	elementPosition := entity.Position{X: x, Y: y}
	//ValidatePosition
	//Validate the position with the service
	if ok := s.service.IsAValidElementPosition(elementPosition); !ok {
		//Invalid position
		return app.ErrorInvalidPosition
	}
	//Return error if exist or nil
	return s.repo.SetElement(chunk, elementPosition, element)
}
