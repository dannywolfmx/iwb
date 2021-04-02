package usecase

import (
	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
	"github.com/dannywolfmx/iwb/app/domain/service"
)

type getSessionChunk struct {
	repo         app.ChunkRepository
	servPosition service.PositionService
}

func NewGetSessionChunk(repo app.ChunkRepository, servPosition service.PositionService) *getSessionChunk {
	return &getSessionChunk{
		repo:         repo,
		servPosition: servPosition,
	}
}

func (g *getSessionChunk) Execute(session *entity.Session) (*entity.Chunk, error) {
	ok := g.servPosition.IsAValidUserPosition(session.UserPosition)

	if !ok {
		return nil, app.ErrorInvalidPosition
	}
	chunkPosition := g.servPosition.CalculateChunkPosition(session.UserPosition)
	//Call the repository to get a chunk in the position
	return g.repo.Get(chunkPosition)
}
