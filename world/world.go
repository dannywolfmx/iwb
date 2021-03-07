package world

type Position struct {
	X, Y uint8
}

type PersistantWorld interface {
	World
	Persist() error
}

type World interface {
	SetPosition(position Position)
	GetPosition() Position
	GetChunk(position Position) *Chunk
}
