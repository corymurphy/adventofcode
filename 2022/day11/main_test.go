package main

import (
	"testing"

	shared "github.com/corymurphy/adventofcode/shared"
)

func Test_Part1(t *testing.T) {
	expected := 10605
	input := shared.ReadInput("input_test")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

// func Test_Part2(t *testing.T) {
// 	expected := 13140
// 	input := shared.ReadInput("input_test")
// 	actual := part2(input)
// 	shared.AssertEqual(t, expected, actual)
// }

// func Test_Part1_Completed(t *testing.T) {
// 	expected := 17840
// 	input := shared.ReadInput("input")
// 	actual := part1(input)
// 	shared.AssertEqual(t, expected, actual)
// }

// func Test_Part2_Completed(t *testing.T) {
// 	expected := 470596
// 	input := shared.ReadInput("input")
// 	actual := part2(input)
// 	shared.AssertEqual(t, expected, actual)
// }
