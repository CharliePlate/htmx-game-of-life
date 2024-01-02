package main

import (
	"context"
	"strconv"
	"strings"

	"githib.com/charlieplate/htmx_todo/components"
	"githib.com/charlieplate/htmx_todo/shared/board"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.File("/static/css/styles.css", "static/css/styles.css")
	e.File("/static/js/index.js", "static/js/index.js")

	e.GET("/", func(c echo.Context) error {
		err := components.Index().Render(context.Background(), c.Response())
		if err != nil {
			return err
		}

		return nil
	})

	e.POST("/board", func(c echo.Context) error {
		cells, err := c.FormParams()
		if err != nil {
			println(err.Error())
			return err
		}

		size := strings.Split(c.FormValue("board-size"), "x")

		height, err := strconv.Atoi(size[0])
		if err != nil {
			println(err.Error())
			return err
		}

		width, err := strconv.Atoi(size[1])
		if err != nil {
			println(err.Error())
			return err
		}

		b := board.MakeEmptyBoard(width, height)

		for k, cell := range cells {
			if k == "board-size" {
				continue
			}

			c, err := board.CellFromJson(cell[0])
			if err != nil {
				println(err.Error())
			}

			b.Cells[c.X()][c.Y()] = c
			if c.Alive {
				b.Cells[c.X()][c.Y()].Revive()
			}
		}

		newBoard := b.PlayLife()

		err = components.Board(newBoard).Render(context.Background(), c.Response())
		if err != nil {
			return err
		}

		return nil
	})

	e.POST("/cell", func(c echo.Context) error {
		params, err := c.FormParams()
		if err != nil {
			println(err.Error())
		}

		var cellStr string
		for _, v := range params {
			cellStr = v[0]
		}

		cell, err := board.CellFromJson(cellStr)

		err = components.Cell(cell.X(), cell.Y(), !cell.Alive).Render(context.Background(), c.Response())
		if err != nil {
			println(err.Error())
			return err
		}

		return nil
	})

	e.Logger.Fatal(e.Start(":3000"))
}
