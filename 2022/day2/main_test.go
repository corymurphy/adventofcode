package main

import (
	"fmt"
	"testing"

	"github.com/corymurphy/adventofcode/shared"
)

func Test_Part1(t *testing.T) {
	expected := 15

	input := shared.ReadInput("input_test")

	actual := part1(input)
	if expected != actual {
		t.Errorf("expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}

func Test_Part2(t *testing.T) {
	expected := 12
	input := shared.ReadInput("input_test")
	actual := part2(input)
	if expected != actual {
		t.Errorf("expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}
