package board

import "encoding/json"

var _NEIGHBORS = [8][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

type Cord struct {
	First  int
	Second int
}

type Board struct {
	Cells  [][]Cell
	Height int
	Width  int
}

type Cell struct {
	Cords [2]int `json:"cords"`
	Alive bool   `json:"alive"`
}

func CellFromJson(s string) (Cell, error) {
	var c Cell
	err := json.Unmarshal([]byte(s), &c)
	if err != nil {
		return Cell{}, err
	}

	return c, nil
}

func (b *Board) PrintBoard() {
	for _, row := range b.Cells {
		println()
		for _, cell := range row {
			if cell.Alive {
				print("O")
			} else {
				print("-")
			}
		}
	}
}

func NewCell(x int, y int, alive bool) Cell {
	return Cell{
		Cords: [2]int{x, y},
		Alive: alive,
	}
}

func (c *Cell) X() int {
	return c.Cords[0]
}

func (c *Cell) Y() int {
	return c.Cords[1]
}

func (c *Cell) Toggle() {
	c.Alive = !c.Alive
}

func (b *Board) MakeStartingBoard() {
	b.Cells = [][]Cell{}
	for i := 0; i < b.Height; i++ {
		b.Cells = append(b.Cells, []Cell{})
		for j := 0; j < b.Width; j++ {
			b.Cells[i] = append(b.Cells[i], NewCell(i, j, false))
		}
	}
}
