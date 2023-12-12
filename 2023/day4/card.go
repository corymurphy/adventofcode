package main

import (
	"slices"
	"strings"

	"github.com/corymurphy/adventofcode/shared"
)

type Card struct {
	winners   []int
	mine      []int
	matching  []int
	instances int
}

func NewCard(value string) *Card {
	valuesSplit := strings.Split(value, " | ")
	winnersUnparsed := strings.Split(valuesSplit[0], " ")
	mineUnparsed := strings.Split(valuesSplit[1], " ")

	winners := []int{}
	allMine := []int{}
	matches := []int{}

	for _, winnerUnparsed := range winnersUnparsed {

		if winnerUnparsed == "" {
			continue
		}
		winners = append(winners, shared.ToInt(winnerUnparsed))
	}

	for _, unparsed := range mineUnparsed {
		if unparsed == "" {
			continue
		}
		allMine = append(allMine, shared.ToInt(unparsed))
	}

	for _, mine := range allMine {
		if slices.Contains(winners, mine) {
			matches = append(matches, mine)
		}
	}

	return &Card{winners: winners, mine: allMine, matching: matches, instances: 1}
}

func (c *Card) AddInstance() {
	c.instances = c.instances + 1
}

func ParseCards(input []string) *[]Card {
	cards := []Card{}

	for _, row := range input {
		cards = append(cards, *NewCard(strings.Split(row, ": ")[1]))
	}

	return &cards
}

func (c *Card) Points() int {
	points := 0
	for range c.matching {
		if points == 0 {
			points = 1
			continue
		}

		points = points * 2

	}

	return points
}
