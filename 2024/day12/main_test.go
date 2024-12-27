package main

import (
	"testing"

	shared "github.com/corymurphy/adventofcode/shared"
)

func Test_Part1_SampleSmall(t *testing.T) {
	var expected int = 140
	input := shared.ReadInput("input_sample_small")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part1_SampleMedium(t *testing.T) {
	var expected int = 772
	input := shared.ReadInput("input_sample_medium")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part1_Sample(t *testing.T) {
	var expected int = 1930
	input := shared.ReadInput("input_sample")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part1(t *testing.T) {
	var expected int = 1304764
	input := shared.ReadInput("input")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

// func Test_Part2_Sample(t *testing.T) {
// 	expected := 1206
// 	input := shared.ReadInput("input_sample")
// 	actual := part2(input)
// 	shared.AssertEqual(t, expected, actual)
// }

// func Test_Part2_SampleMedium(t *testing.T) {
// 	expected := 368
// 	input := shared.ReadInput("input_sample_medium")
// 	actual := part2(input)
// 	shared.AssertEqual(t, expected, actual)
// }

func Test_Part2_SampleMedium2(t *testing.T) {
	expected := 236
	input := shared.ReadInput("input_sample_medium2")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}

// func Test_Part2_SampleSmall(t *testing.T) {
// 	expected := 80
// 	input := shared.ReadInput("input_sample_small")
// 	actual := part2(input)
// 	shared.AssertEqual(t, expected, actual)
// }

// func Test_Part2(t *testing.T) {
// 	expected := 0 // 575843 too low
// 	input := shared.ReadInput("input")
// 	actual := part2(input)
// 	shared.AssertEqual(t, expected, actual)
// }
