package board

import "testing"

func TestCellFromJson(t *testing.T) {
	c, err := CellFromJson(`{"cords":[0,0],"alive":true}`)
	if err != nil {
		t.Error(err)
	}

	if c.X() != 0 {
		t.Errorf("Expected x to be 0, got %d", c.X())
	}
	if c.Y() != 0 {
		t.Errorf("Expected y to be 0, got %d", c.Y())
	}

	if c.Alive != true {
		t.Errorf("Expected alive to be true, got %t", c.Alive)
	}
}

func TestNewCell(t *testing.T) {
	c := NewCell(0, 0, true)

	if c.X() != 0 {
		t.Errorf("Expected x to be 0, got %d", c.X())
	}

	if c.Y() != 0 {
		t.Errorf("Expected y to be 0, got %d", c.Y())
	}

	if c.Alive != true {
		t.Errorf("Expected alive to be true, got %t", c.Alive)
	}
}

func TestCellToggle(t *testing.T) {
	c := NewCell(0, 0, true)

	if c.Alive != true {
		t.Errorf("Expected alive to be true, got %t", c.Alive)
	}

	c.Toggle()

	if c.Alive != false {
		t.Errorf("Expected alive to be false, got %t", c.Alive)
	}
}

func TestAliveNeighbors(t *testing.T) {
	b := MakeEmptyBoard(3, 3)

	b.Cells[0][0].Toggle()
	b.Cells[0][1].Toggle()
	b.Cells[0][2].Toggle()

	aliveNeighbors := b.AliveNeighbors(&b.Cells[1][1])

	if aliveNeighbors != 3 {
		t.Errorf("Expected alive neighbors to be 3, got %d", aliveNeighbors)
	}
}

func AliveWithThree(t *testing.T) {
	b := MakeEmptyBoard(3, 3)

	b.Cells[0][0].Toggle()
	b.Cells[0][1].Toggle()
	b.Cells[0][2].Toggle()

	newBoard := MakeEmptyBoard(3, 3)

	b.ApplyRules(&b.Cells[1][1], b.AliveNeighbors(&b.Cells[1][1]), &newBoard)

	if newBoard.Cells[1][1].Alive != true {
		t.Errorf("Expected cell to be alive, got %t", newBoard.Cells[1][1].Alive)
	}
}

func DeadWithFour(t *testing.T) {
	b := MakeEmptyBoard(3, 3)

	b.Cells[0][0].Toggle()
	b.Cells[0][1].Toggle()
	b.Cells[0][2].Toggle()
	b.Cells[1][0].Toggle()

	newBoard := MakeEmptyBoard(3, 3)

	b.ApplyRules(&b.Cells[1][1], b.AliveNeighbors(&b.Cells[1][1]), &newBoard)

	if newBoard.Cells[1][1].Alive != false {
		t.Errorf("Expected cell to be dead, got %t", newBoard.Cells[1][1].Alive)
	}
}
