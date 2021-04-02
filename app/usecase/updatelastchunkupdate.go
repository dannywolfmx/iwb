package usecase

import (
	"time"

	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
)

type updateLastChunkUpdate struct {
	repo app.SessionRepository
}

func NewUpdateLastChunkUpdate(repo app.SessionRepository) *updateLastChunkUpdate {
	return &updateLastChunkUpdate{
		repo: repo,
	}
}

func (u *updateLastChunkUpdate) Execute(session *entity.Session, time time.Time) (*entity.Session, error) {
	session.LastChunkUpdate = time
	err := u.repo.Update(session)
	return session, err
}
