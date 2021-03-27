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
	UsersElements map[ID]UserElements
	LastUpdate    time.Time
}

//NewChunk will return a well formed Chunk reference with a empty map of elements
func NewChunk() *Chunk {
	elements := make(Elements)
	return &Chunk{
		Elements: elements,
	}
}

//Business logic

//NewChunkWithElements will return a well formed Chunk reference with the given elements
func NewChunkWithElements(elements Elements) *Chunk {
	return &Chunk{
		Elements: elements,
	}
}

//GetElement return a Element in the given position if extist, or return a space like element if the element doesn't exist
func (c *Chunk) GetElement(position Position) Element {
	if element, ok := c.Elements[position]; ok {
		return element
	}
	//Return a space like a new rune
	return ' '
}

//GetElements return all the avaible elements in the chunk
func (c *Chunk) GetElements() Elements {
	return c.Elements
}

//SetElement will set a element in the given position, if the is already in use, the new element will overwrite the old element
//The position need to be a valid one
func (c *Chunk) SetElement(position Position, element Element) {
	c.Elements[position] = element
}
