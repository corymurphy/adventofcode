package main

import (
	"testing"

	shared "github.com/corymurphy/adventofcode/shared"
)

func Test_Part1_Sample(t *testing.T) {
	expected := 7036
	input := shared.ReadInput("input_sample")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}
func Test_Part1_Sample2(t *testing.T) {
	expected := 11048
	input := shared.ReadInput("input_sample2")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part1(t *testing.T) {
	expected := 92432
	input := shared.ReadInput("input")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2_Sample(t *testing.T) {
	expected := 45
	input := shared.ReadInput("input_sample")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2_Sample2(t *testing.T) {
	expected := 64
	input := shared.ReadInput("input_sample2")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2(t *testing.T) {
	expected := 458
	input := shared.ReadInput("input")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}
