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
	game := NewGame(input, 3)
	result := game.Play(20)
	return result
}

func part2(input []string) int {
	game := NewGame(input, 1)
	result := game.Play(10000)
	return result
}
