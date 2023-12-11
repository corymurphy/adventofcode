package main

import (
	"strings"

	"github.com/corymurphy/adventofcode/shared"
)

type Game struct {
	id     int
	rounds *[]Round
}

func NewGame(results string) *Game {

	game := &Game{}

	gameInitial := strings.Split(results, ": ")

	game.id = shared.ToInt(gameInitial[0][5:])
	game.rounds = ParseRounds(gameInitial[1])
	return game
}

func ParseGames(results []string) *[]Game {

	games := []Game{}

	for _, result := range results {
		games = append(games, *NewGame(result))
	}

	return &games
}

func (g *Game) IsPossible(requirements map[string]int) bool {
	for _, round := range *g.rounds {
		for cube, requiredCount := range requirements {
			if revealCount, contains := round.revealed[cube]; contains {
				if requiredCount < revealCount {
					return false
				}
			}
		}

	}

	return true
}

func (g *Game) MinimumPossible() map[string]int {
	min := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, round := range *g.rounds {
		for cube, revealCount := range round.revealed {
			if min[cube] < revealCount {
				min[cube] = revealCount
			}
		}
	}
	return min
}
