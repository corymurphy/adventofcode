package main

import (
	"fmt"
	"strconv"
	"strings"
)

func NewBoards(input []string) []Board {

	boards := make([]Board, 0)

	for i := 0; i < len(input); i++ {

		row := input[i]

		if i == 0 {
			fmt.Println("ignoring first line")
			continue
		}

		if len(strings.TrimSpace(row)) == 0 {
			continue // empty line
		}

		if len(strings.TrimSpace(row)) > 0 {
			board := NewBoard(input[i : i+5])
			boards = append(boards, board)
		}
		i = i + 5
	}
	return boards
}

type Cell struct {
	Value string
	Hit   bool
}

type Board struct {
	Values [][]Cell
	Winner bool
}

func NewBoard(input []string) Board {

	cols := make([][]Cell, len(input))

	for col := 0; col < len(input); col++ {

		rows := make([]Cell, 5)

		elementIndex := 0

		for cell := 0; cell < 5 && elementIndex < len(input[col]); cell++ {

			rows[cell] = Cell{
				Hit:   false,
				Value: strings.TrimSpace(input[col][elementIndex : elementIndex+2]),
			}

			elementIndex = elementIndex + 3
		}
		cols[col] = rows
	}
	return Board{
		Values: cols,
		Winner: false,
	}
}

func (b *Board) ToString() {
	for _, col := range b.Values {
		fmt.Println(col)
	}
}

func (b *Board) Play(drawn string) {
	for row, rows := range b.Values {
		for i, cell := range rows {
			if cell.Value == drawn {
				cell.Hit = true
				b.Values[row][i] = cell
			}
		}
	}
}

func horizontalWinner(row []Cell, colIndex int) bool {

	if colIndex > 0 {
		return false
	}

	for _, cell := range row {
		if !cell.Hit {
			return false
		}
	}
	return true
}

func verticalWinner(board [][]Cell, colIndex int, rowIndex int) bool {
	if rowIndex > 0 {
		return false
	}

	hits := 0
	for col, cell := range board[rowIndex] {

		hits = 0

		if !cell.Hit {
			continue
		}

		if cell.Hit {
			for i := 0; i < len(board[rowIndex]); i++ {

				if !board[i][col].Hit {
					break
				} else {
					hits++
				}
			}

		}

		if hits == len(board[0]) {
			return true
		}
	}

	return false

}

func (b *Board) IsWinner() bool {
	for rowIndex, cells := range b.Values {
		for colIndex, _ := range cells {

			if horizontalWinner(cells, colIndex) {
				b.Winner = true
				return true
			}

			if verticalWinner(b.Values, colIndex, rowIndex) {
				b.Winner = true
				return true
			}
		}
	}
	return false
}

func (b *Board) SumUnmarked() int {
	sum := 0

	for _, row := range b.Values {
		for _, cell := range row {
			if !cell.Hit {
				value, _ := strconv.ParseInt(cell.Value, 10, 32)
				sum = sum + int(value)
			}
		}
	}

	return sum
}
