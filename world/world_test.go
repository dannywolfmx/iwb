package world

import "testing"

func TestChunk(t *testing.T) {
	chunk := NewChunk()
	if chunk.GetRune(4, 4) != ' ' {
		t.Fail()
	}
	chunk.SetRune(4, 4, 'k')
	if chunk.GetRune(4, 4) != 'k' {
		t.Fail()
	}
}

func TestWorld(t *testing.T) {
	world := NewWorld()
	chunk := world.GetChunk(5, 5)
	chunk.SetRune(4, 4, 'k')

	chunk2 := world.GetChunk(5, 5)
	if chunk != chunk2 {
		t.Fail()
	}
	if chunk.GetRune(4, 4) != 'k' {
		t.Fail()
	}
}

func TestGetChunkAtPos(t *testing.T) {
	var x, y int
	x, y = GetChunkAtPos(20, 20)
	if x != 0 || y != 0 {
		t.Fail()
	}

	x, y = GetChunkAtPos(1000, 20)
	if x != 3 || y != 0 {
		t.Fail()
	}

	x, y = GetChunkAtPos(1000, 1000)
	if x != 3 || y != 3 {
		t.Fail()
	}
}
