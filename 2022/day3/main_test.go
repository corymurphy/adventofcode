package main

import (
	"fmt"
	"testing"

	"github.com/corymurphy/adventofcode/shared"
)

func Test_Part1(t *testing.T) {
	expected := 157

	input := shared.ReadInput("input_test")

	actual := part1(input)
	if expected != actual {
		t.Errorf("expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}

func Test_Part2(t *testing.T) {
	expected := 70
	input := shared.ReadInput("input_test")
	actual := part2(input)
	if expected != actual {
		t.Errorf("expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}

func Test_Part1_Completed(t *testing.T) {
	expected := 7727

	input := shared.ReadInput("input")

	actual := part1(input)
	if expected != actual {
		t.Errorf("expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}

func Test_Part2_Completed(t *testing.T) {
	expected := 2609

	input := shared.ReadInput("input")

	actual := part1(input)
	if expected != actual {
		t.Errorf("expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}
