package usecase

import (
	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
)

type generateTokenFuncion = func(user *entity.User) (string, error)

type generateSession struct {
	repo              app.SessionRepository
	generateTokenFunc generateTokenFuncion
}

//NewGenerateSession constructor to generate an filled "generateSession"
func NewGenerateSession(repo app.SessionRepository, generateTokenFunc generateTokenFuncion) *generateSession {
	return &generateSession{
		repo:              repo,
		generateTokenFunc: generateTokenFunc,
	}
}

//Execute Generate a new user session to send to the client
func (g *generateSession) Execute(user *entity.User) (*entity.Session, error) {
	//Generate token
	token, err := g.generateTokenFunc(user)

	if err != nil {
		return nil, err
	}

	//generate a new session entity
	session := entity.NewSession(user.Name, token)
	//Save new session in the repository
	err = g.repo.Save(session)
	//return session
	return session, err
}
