//List of custom errors to the application
package app

import "errors"

var (
	//GetChunk errors
	ErrorInvalidPosition  = errors.New("Error: Invalid position")
	ErrorOnGetChunkFromDB = errors.New("Error: repository get chunk")

	//SetElement erros
	ErrorOnSetElementToDB = errors.New("Error: repository set element")

	//User errors
	ErrorUserAlreadyExist = errors.New("Error: user already exist")
)
