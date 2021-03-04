package world

import "testing"

type ChunkTest struct {
	testRune rune
}

func (c *ChunkTest) GetRune(position Position) rune {
	return c.testRune
}
func (c *ChunkTest) SetRune(position Position, char rune) {
	c.testRune = char
}
func (c *ChunkTest) LastUpdatedAt() int64 {
	return 0
}

func TestChunk(t *testing.T) {
	chunk := &ChunkTest{}
	position := Position{4, 4}
	chunk.SetRune(position, 'k')
	if chunk.GetRune(position) != 'k' {
		t.Fatal("Should get a rune in the position")
	}
}
