//List of custom errors to the application
package app

import "errors"

var (
	//GetChunk errors
	ErrorInvalidPosition  = errors.New("Invalid position")
	ErrorOnGetChunkFromDB = errors.New("Error repository get chunk")

	//SetElement erros
	ErrorOnSetElementToDB = errors.New("Error repository set element")

	//
)
