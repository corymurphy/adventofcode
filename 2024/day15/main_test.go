package main

import (
	"testing"

	shared "github.com/corymurphy/adventofcode/shared"
)

func Test_Part1_SampleSmall(t *testing.T) {
	expected := 2028
	input := shared.ReadInput("input_sample_small")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part1_Sample(t *testing.T) {
	expected := 10092
	input := shared.ReadInput("input_sample")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part1(t *testing.T) {
	expected := 1360570
	input := shared.ReadInput("input")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2_Sample(t *testing.T) {
	expected := 9021
	input := shared.ReadInput("input_sample")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2(t *testing.T) {
	expected := 0
	input := shared.ReadInput("input")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}
