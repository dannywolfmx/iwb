package mock

import (
	"github.com/dannywolfmx/iwb/app"
	"github.com/dannywolfmx/iwb/app/domain/entity"
)

type UserRepository struct {
	user     *entity.User
	position entity.Position
}

func (u *UserRepository) SetPosition(user *entity.User, position entity.Position) error {
	if user == u.user && position == u.position {
		return nil
	}

	return app.ErrorInvalidPosition
}
