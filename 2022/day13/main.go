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
	ordered := 0

	signal := NewSignal(input)

	for _, pair := range *signal {
		if pair.InOrder() {
			ordered = ordered + pair.index
		}
	}
	return ordered
}

func part2(input []string) int {
	return 0
}
