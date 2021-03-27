package entity

import "time"

type Token = string

type Session struct {
	//ActualChunkPosition is the position of the session in the world
	ActualChunkPosition Position

	//Account information
	Account *User

	//To Check if the session is outdated
	LastChunkUpdate time.Time

	//UID to know the session request from the user
	Token Token
}
