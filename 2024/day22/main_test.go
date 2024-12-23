package main

import (
	"testing"

	shared "github.com/corymurphy/adventofcode/shared"
)

func Test_Part1_Sample(t *testing.T) {
	expected := 37327623
	input := shared.ReadInput("input_sample")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part1(t *testing.T) {
	expected := 16953639210
	input := shared.ReadInput("input")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2_Sample(t *testing.T) {
	expected := 23
	input := shared.ReadInput("input_sample2")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2(t *testing.T) {
	expected := 0
	input := shared.ReadInput("input")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}
