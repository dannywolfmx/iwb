package usecase

import (
	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
	"github.com/dannywolfmx/iwb/app/domain/service"
)

type generateSession struct {
	repo            app.SessionRepository
	servicePosition *service.PositionService
	serviceToken    *service.TokenService
}

func NewGenerateSession(repo app.SessionRepository, servPosition *service.PositionService, servToken *service.TokenService) *generateSession {
	return &generateSession{
		repo:            repo,
		servicePosition: servPosition,
		serviceToken:    servToken,
	}
}

func (g *generateSession) Execute(user *entity.User) (*entity.Session, error) {
	//Generate token
	token, err := g.serviceToken.GenerateJWT(user)

	if err != nil {
		return nil, err
	}

	//Get random position
	position := g.servicePosition.GetDefaultPosition()
	//generate a new session entity
	session := entity.NewSession(user.Name, token, position)
	//Save new session in the repository
	err = g.repo.Save(session)
	//return session
	return session, err
}
