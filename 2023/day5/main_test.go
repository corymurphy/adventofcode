package main

import (
	"testing"

	shared "github.com/corymurphy/adventofcode/shared"
)

func Test_Part1TestInput(t *testing.T) {
	expected := 13
	input := shared.ReadInput("input_test")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part1Input(t *testing.T) {
	expected := 23941
	input := shared.ReadInput("input")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2TestInput(t *testing.T) {
	expected := 30
	input := shared.ReadInput("input_test")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2Input(t *testing.T) {
	expected := 5571760
	input := shared.ReadInput("input")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}
