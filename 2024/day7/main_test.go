package main

import (
	"testing"

	shared "github.com/corymurphy/adventofcode/shared"
)

// func Test_Part1_Sample(t *testing.T) {
// 	var expected uint64 = 3749
// 	input := shared.ReadInput("input_sample")
// 	actual := part1(input)
// 	shared.AssertEqual(t, expected, actual)
// }

// func Test_Part1(t *testing.T) {
// 	var expected uint64 = 6392012777720
// 	input := shared.ReadInput("input")
// 	actual := part1(input)
// 	shared.AssertEqual(t, expected, actual)
// }

func Test_Part2_Sample(t *testing.T) {
	var expected uint64 = 11387
	input := shared.ReadInput("input_sample")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2(t *testing.T) {
	expected := 0
	// 31056542151082 too low
	input := shared.ReadInput("input")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}
