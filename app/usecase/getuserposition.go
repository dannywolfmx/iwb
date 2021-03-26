package usecase

import (
	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
)

type getUserPosition struct {
	repo app.UserRepository
}

func NewGetUserPosition(repo app.UserRepository) *getUserPosition {
	return &getUserPosition{
		repo: repo,
	}
}

//Execute the usecase
func (g *getUserPosition) Execute(user *entity.User) (entity.Position, entity.Position, error) {
	return g.repo.GetPosition(user)
}
