package app

import (
	"github.com/dannywolfmx/iwb/app/domain/entity"
)

//go:generate mockgen --destination=./repository/mock/session.go . SessionRepository
type SessionRepository interface {
	Save(session *entity.Session) error
	Update(session *entity.Session) error
}

//go:generate mockgen --destination=./repository/mock/user.go . UserRepository
type UserRepository interface {
	Exist(user *entity.User) bool
}

//go:generate mockgen --destination=./repository/mock/chunk.go . ChunkRepository
type ChunkRepository interface {
	Get(position entity.Position) (*entity.Chunk, error)
}
