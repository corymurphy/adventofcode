package main

import (
	"fmt"

	"github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(directions []string) int {

	result := 0

	for _, direction := range directions[0] {
		if direction == '(' {
			result++
		} else if direction == ')' {
			result--
		}
	}
	return result
}

func part2(directions []string) int {

	floor := 0

	for i, direction := range directions[0] {
		if direction == '(' {
			floor++
		} else if direction == ')' {
			floor--
		}

		if floor == -1 {
			return i + 1
		}
	}
	return floor
}
