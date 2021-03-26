package usecase

import (
	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
	"github.com/dannywolfmx/iwb/app/domain/service"
)

type setUserPosition struct {
	repo            app.UserRepository
	servicePosition *service.PositionService
	serviceUser     *service.UserService
}

//NewSetUserPosition return a setUserPosition reference
func NewSetUserPosition(repo app.UserRepository) *setUserPosition {
	return &setUserPosition{
		repo:            repo,
		servicePosition: &service.PositionService{},
		serviceUser: &service.UserService{
			Repo: repo,
		},
	}
}

func (s *setUserPosition) Execute(user *entity.User, position entity.Position) error {
	//Validate the position
	if ok := s.servicePosition.IsAValidUserPosition(position); !ok {
		return app.ErrorInvalidPosition
	}

	//Check if is a valid user
	if ok := s.serviceUser.IsAValidUser(user); !ok {
		return app.ErrorInvalidPosition
	}

	//Set value and return an error if exist
	return s.repo.SetPosition(user, position)
}
