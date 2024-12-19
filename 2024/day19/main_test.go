package main

import (
	"testing"

	shared "github.com/corymurphy/adventofcode/shared"
)

func Test_Part1_Sample(t *testing.T) {
	expected := 6
	input := shared.ReadInput("input_sample")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part1(t *testing.T) {
	// 243 too low
	// 257 too low
	// 264 too high
	// 257 wrong
	// 252 wrong
	// 255 wrong

	// untest
	// 256
	// 255
	expected := 260
	input := shared.ReadInput("input")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2_Sample(t *testing.T) {
	expected := 0
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
