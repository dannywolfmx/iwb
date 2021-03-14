package entity

type PersistantWorld interface {
	World
	Persist() error
}

type World interface {
	SetPosition(viewport, chunkLocation Position)
	//GetPosition return the viewport, and chunk location
	GetPosition() (Position, Position)
	GetChunk(position Position) *Chunk
}
