package main

import (
	"testing"

	shared "github.com/corymurphy/adventofcode/shared"
)

func Test_Part1(t *testing.T) {
	expected := 24000
	input := shared.ReadInput("input_test")
	actual := mostCarriedCalories(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2(t *testing.T) {
	expected := 45000
	input := shared.ReadInput("input_test")
	actual := topThreeCarriedCalories(input)
	shared.AssertEqual(t, expected, actual)
}
