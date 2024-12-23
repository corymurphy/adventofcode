package main

import (
	"fmt"

	shared "github.com/corymurphy/adventofcode/shared"
)

// <A^A>^^AvvvA
// v<<A>>^A <A>A vA<^AA>A<vAAA>^A

// < --- v<<A>>^A
// A
type DirectionalKeyPad struct{}

func (d *DirectionalKeyPad) Down() (seq []string) {
	return []string{"<"}
}

func (d *DirectionalKeyPad) Left() (seq []string) {
	return []string{"v", "<", "<", "A"}
}

func main() {
	input := shared.ReadInput("input_sample")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) (answer int) {

	return answer
}

func part2(input []string) (answer int) {

	return answer
}
