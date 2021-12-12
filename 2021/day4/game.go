package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	Boards []Board
	Drawn  []string
}

func NewGame(input []string) Game {
	game := Game{
		Boards: NewBoards(input),
		Drawn:  parseDrawn(input),
	}
	return game
}

func parseDrawn(input []string) []string {
	return strings.Split(input[0], ",") // TODO: add defense
}

func (g *Game) ToString() {

	for _, b := range g.Boards {
		fmt.Println("")
		b.ToString()
		fmt.Println("")
	}
}

func (g *Game) PlayPart1() int {
	for _, drawn := range g.Drawn {
		for _, board := range g.Boards {
			board.Play(drawn)
			if board.IsWinner() {
				board.ToString()
				drawnValue, _ := strconv.ParseInt(drawn, 10, 32)
				return int(drawnValue) * board.SumUnmarked()
			}
		}
	}
	return 0
}

func (g *Game) AllBoardsWon() bool {

	for _, board := range g.Boards {
		if !board.IsWinner() { // not efficient but for some reason, board.Winner isn't storing the value
			return false
		}
	}
	return true
}

func (g *Game) PlayPart2() int {

	for _, drawn := range g.Drawn {
		for _, board := range g.Boards {
			board.Play(drawn)
			if board.IsWinner() && g.AllBoardsWon() {
				board.ToString()
				drawnValue, _ := strconv.ParseInt(drawn, 10, 32)
				return int(drawnValue) * board.SumUnmarked()
			}
		}
	}

	return 0

}
