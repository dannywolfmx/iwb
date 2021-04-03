package entity

import "time"

type Elements = map[Position]Element

type UserElements struct {
	Account  *User
	Elements Elements
}

type Chunk struct {
	//map of users elements in the chunk
	//ID is the user id
	UsersElements map[ID]*UserElements
	LastUpdate    time.Time
}

//Create a empty chunk
func NewChunk() *Chunk {
	return &Chunk{
		UsersElements: make(map[ID]*UserElements),
		LastUpdate:    time.Now(),
	}
}
