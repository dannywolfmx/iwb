//usecase.go descripbe the interface of the "app" usecases
package app

import "github.com/dannywolfmx/iwb/app/domain/entity"

//GenerateSession will generate a session to keep tracking of the user in the rest of the application
//It is a friendly way to request a loggin to the user, just a name
type GenerateSession interface {
	Execute(user *entity.User) (*entity.Session, error)
}

// TODO implement
//
//GetSessionChunkElements
//Return the actual session's chunk
type GetSessionChunk interface {
	Execute(session *entity.Session) (*entity.Chunk, error)
}

// TODO implement
//
//SetElement
//Set an element in the actual position of the user
//And move the user to the next position
type SetElement interface {
	Execute(session *entity.Session, element entity.Element) error
}

// TODO implement
//
//GetCurrentPosition
//Get the position of the user
type GetCurrentPosition interface {
	Execute(session *entity.Session) (entity.Position, error)
}

//SetCurrentPosition
//Set a position of a current user into the actual world chunk
//You can use this interface to move a user in a chunk
type SetCurrentPosition interface {
	//Execute the usecase
	Execute(session *entity.Session, position entity.Position) error
}

// TODO implement
// NeedSessionAChunkUpdate check if the session chunk is outdated
type NeedSessionAChunkUpdate interface {
	Execute(session *entity.Session) (bool, error)
}
