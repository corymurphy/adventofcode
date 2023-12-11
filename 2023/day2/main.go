package main

import (
	"fmt"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) int {

	requirements := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	possibleSum := 0
	games := *ParseGames(input)

	for _, game := range games {
		if game.IsPossible(requirements) {
			possibleSum = possibleSum + game.id
		}
	}
	return possibleSum
}

func part2(input []string) int {

	games := *ParseGames(input)

	powerSum := 0
	for _, game := range games {
		min := game.MinimumPossible()

		power := min["blue"] * min["green"] * min["red"]
		powerSum = powerSum + power
	}
	return powerSum
}
