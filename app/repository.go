package app

import "github.com/dannywolfmx/iwb/app/domain/entity"

//go:generate mockgen --destination=./repository/mock/session.go . SessionRepository
type SessionRepository interface {
	Save(*entity.Session) error
}

//go:generate mockgen --destination=./repository/mock/user.go . UserRepository
type UserRepository interface {
	Exist(user *entity.User) bool
}
