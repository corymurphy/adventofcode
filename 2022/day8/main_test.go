package main

import (
	"testing"

	"github.com/corymurphy/adventofcode/shared"
)

func Test_Part1(t *testing.T) {
	expected := 21
	input := shared.ReadInput("input_test")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_GetLargest(t *testing.T) {
	expected := 5
	input := []int{1, 2, 5, 3, 4}
	actual := getLargest(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2(t *testing.T) {
	expected := 8
	input := shared.ReadInput("input_test")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part1_Completed(t *testing.T) {
	expected := 1700
	input := shared.ReadInput("input")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2_Completed(t *testing.T) {
	expected := 470596
	input := shared.ReadInput("input")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}
