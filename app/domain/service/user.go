package service

import (
	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
)

type UserService struct {
	Repo app.UserRepository
}

func (s *UserService) IsAValidUser(user *entity.User) bool {
	//If not error then the user exist
	return s.Repo.Exist(user)
}
