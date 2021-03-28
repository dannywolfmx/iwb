package app

import "github.com/dannywolfmx/iwb/app/domain/entity"

type SessionRepository interface {
	Save(*entity.Session) error
}
