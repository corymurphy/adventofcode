package main

import (
	"fmt"
)

func main() {
	input := readInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) int {
	game := NewGame(input)
	return game.PlayPart1()
}

func part2(input []string) int {
	game := NewGame(input)
	return game.PlayPart2()
}
