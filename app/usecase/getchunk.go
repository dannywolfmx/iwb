package usecase

import (
	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
	"github.com/dannywolfmx/iwb/app/domain/service"
)

type getChunk struct {
	repo    app.WorldRepository
	service *service.PositionService
}

//NewWorldUsecase set a repo
func NewGetChunk(repo app.WorldRepository) *getChunk {
	return &getChunk{
		repo:    repo,
		service: &service.PositionService{},
	}
}

//Execute will performe the GetCHunk usecase
//pass a valid x, y postion and return a chunk pointer or error
func (g *getChunk) Execute(position entity.Position) (*entity.Chunk, error) {
	//Validate the position with the service
	if ok := g.service.IsAValidChunkPosition(position); !ok {
		//Invalid position
		return nil, app.ErrorInvalidPosition
	}

	//Get Chunk in the position
	chunk, err := g.repo.GetChunk(position)

	if err != nil {
		return nil, err
	}

	//Return result
	return chunk, nil
}
