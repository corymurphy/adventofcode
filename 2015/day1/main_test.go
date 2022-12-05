package main

import (
	"testing"

	"github.com/corymurphy/adventofcode/shared"
)

func Test_Part1(t *testing.T) {
	expected := -3

	input := shared.ReadInput("input_test")

	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2(t *testing.T) {
	expected := 5
	input := []string{"()())"}
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}
