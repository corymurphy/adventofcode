package main

import (
	"fmt"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %s\n\n", part1)
	fmt.Printf("\nPart 2 answer: %s\n\n", part2)
}

func part1(input []string) string {
	dock := NewDock(input)
	dock.Rearrange()
	return dock.TopCrates()
}

func part2(input []string) string {
	dock := NewDock(input)
	dock.RearrangePart2()
	return dock.TopCrates()
}
