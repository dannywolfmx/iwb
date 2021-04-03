package service

import (
	"github.com/dannywolfmx/iwb/app/domain/entity"
)

type ChunkService struct{}

func (c *ChunkService) AddElementToUserChunk(id entity.ID, e entity.Element, p entity.Position, chunk *entity.Chunk) {
	_, ok := chunk.UsersElements[id]

	//https://www.programmersought.com/article/92415522009/
	if !ok {
		chunk.UsersElements[id] = &entity.UserElements{
			Account:  nil,
			Elements: make(map[entity.Position]rune),
		}
	}
	chunk.UsersElements[id].Elements[p] = e
}
