//This packege is made to implement the repository interface, to test propose
package mock

import (
	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
)

type WorldMockRepository struct {
	Position entity.Position
	Chunk    *entity.Chunk
	Element  entity.Element
}

func (t *WorldMockRepository) GetChunk(position entity.Position) (*entity.Chunk, error) {
	//Check the position to be the same
	if t.Position == position {
		return entity.NewChunk(), nil
	}
	//Return nil as chunk to represent an error in the validate
	return nil, app.ErrorInvalidPosition
}

func (t *WorldMockRepository) SetElement(chunk *entity.Chunk, position entity.Position, element entity.Element) error {
	if t.Position != position || t.Element != element {
		return app.ErrorOnSetElementToDB
	}
	return nil
}
