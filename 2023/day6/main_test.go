package main

import (
	"testing"

	shared "github.com/corymurphy/adventofcode/shared"
)

func Test_Part1TestInput(t *testing.T) {
	var expected int64 = 288
	input := shared.ReadInput("input_test")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part1Input(t *testing.T) {
	var expected int64 = 5133600
	input := shared.ReadInput("input")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2TestInput(t *testing.T) {
	var expected int64 = 71503
	input := shared.ReadInput("input_test2")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2Input(t *testing.T) {
	var expected int64 = 40651271
	input := shared.ReadInput("input2")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}
