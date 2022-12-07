package main

import (
	"fmt"
	"testing"

	"github.com/corymurphy/adventofcode/shared"
)

func Test_Part1(t *testing.T) {
	expected := 95437
	input := shared.ReadInput("input_test")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2(t *testing.T) {
	expected := 24933642
	input := shared.ReadInput("input_test")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part1_Completed(t *testing.T) {
	expected := 1084134
	input := shared.ReadInput("input")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2_Completed(t *testing.T) {
	expected := 6183184
	input := shared.ReadInput("input")
	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Tree(t *testing.T) {
	fmt.Println("")
	input := shared.ReadInput("input_test")
	fs := NewFileSystem(input)
	fs.Print()
	fmt.Println()
}
