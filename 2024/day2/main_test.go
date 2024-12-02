package main

import (
	"testing"

	shared "github.com/corymurphy/adventofcode/shared"
)

func Test_Part1_Sample(t *testing.T) {
	expected := 2
	input := shared.ReadInput("input_sample")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part1(t *testing.T) {
	expected := 591
	input := shared.ReadInput("input")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2_Sample(t *testing.T) {
	expected := 4
	input := shared.ReadInput("input_sample")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2(t *testing.T) {
	expected := 621
	input := shared.ReadInput("input")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}
