package main

import (
	"testing"

	shared "github.com/corymurphy/adventofcode/shared"
)

func Test_Part1(t *testing.T) {
	expected := 54990
	input := shared.ReadInput("input")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2(t *testing.T) {
	expected := 54473
	input := shared.ReadInput("input")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Partb(t *testing.T) {
	expected := 73
	input := shared.ReadInput("input_test2b")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Partc(t *testing.T) {
	expected := 73
	input := shared.ReadInput("input_test2c")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}
