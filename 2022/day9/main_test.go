package main

import (
	"testing"

	"github.com/corymurphy/adventofcode/shared"
)

func Test_Part1(t *testing.T) {
	expected := 13
	input := shared.ReadInput("input_test")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

// func Test_Part1_2(t *testing.T) {
// 	expected := 13
// 	input := shared.ReadInput("input_test2")
// 	// actual := part1(input)

// 	grid := NewGrid(30)
// 	grid = SimulateBridge(input, grid)
// 	actual := CountVisited(grid)
// 	PrintGridBackwards(grid)
// 	shared.AssertEqual(t, expected, actual)
// }

// func Test_GetBridgeSize(t *testing.T) {
// 	expected := 6
// 	input := shared.ReadInput("input_test")
// 	actual := GetBridgeSize(input)
// 	shared.AssertEqual(t, expected, actual)
// }

func Test_Part2(t *testing.T) {
	expected := 38
	input := shared.ReadInput("input_test_part2")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part1_Completed(t *testing.T) {
	expected := 6175
	input := shared.ReadInput("input")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

// func Test_Part2_Completed(t *testing.T) {
// 	expected := 470596
// 	input := shared.ReadInput("input")
// 	actual := part2(input)
// 	shared.AssertEqual(t, expected, actual)
// }
