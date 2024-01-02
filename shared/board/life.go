package board

func (b *Board) ApplyRules(c *Cell, aliveNeighbors int, newBoard *Board) {
	x := c.X()
	y := c.Y()
	switch c.Alive {
	case true:
		if aliveNeighbors < 2 || aliveNeighbors > 3 {
			newBoard.Cells[x][y].Kill()
			return
		} else {
			newBoard.Cells[x][y].Revive()
			return
		}
	case false:
		if aliveNeighbors == 3 {
			newBoard.Cells[x][y].Revive()
			return
		}
	}

	newBoard.Cells[x][y] = NewCell(x, y, false)
}

func (b *Board) AliveNeighbors(c *Cell) int {
	aliveNeighbors := 0
	for _, neighbor := range _NEIGHBORS {
		x := c.X() + neighbor[0]
		y := c.Y() + neighbor[1]

		// OOB check
		if x < 0 || x > b.Width-1 || y < 0 || y > b.Height-1 {
			continue
		}

		if b.Cells[x][y].Alive {
			aliveNeighbors += 1
		}
	}

	return aliveNeighbors
}

func MakeEmptyBoard(width int, height int) Board {
	b := Board{
		Height: height,
		Width:  width,
		Cells:  make([][]Cell, height),
	}

	for i := range b.Cells {
		b.Cells[i] = make([]Cell, width)
	}

	return b
}

func (b *Board) PlayLife() Board {
	newBoard := MakeEmptyBoard(b.Width, b.Height)
	for _, row := range b.Cells {
		for _, cell := range row {
			newBoard.Cells[cell.X()][cell.Y()] = cell
			aliveNeighbors := b.AliveNeighbors(&cell)
			newBoard.ApplyRules(&cell, aliveNeighbors, &newBoard)
		}
	}
	return newBoard
}

func (c *Cell) Kill() {
	c.Alive = false
}

func (c *Cell) Revive() {
	c.Alive = true
}
